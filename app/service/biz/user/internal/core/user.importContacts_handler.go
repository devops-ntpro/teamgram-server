/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package core

import (
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"
	"time"
)

type contactItem struct {
	c               *mtproto.InputContact
	unregistered    bool  // не зарегистрирован
	userId          int64 // ID зарегистрированного пользователя
	contactId       int64 // уже зарегистрированный в качестве моего контакта
	importContactId int64 // обратный контакт (0 или userId). Тот, кто добавил меня в качестве контакта
}

// UserImportContacts
// user.importContacts user_id:long contacts:Vector<InputContact> = UserImportedContacts;
func (c *UserCore) UserImportContacts(in *user.TLUserImportContacts) (*user.UserImportedContacts, error) {
	var (
		contacts          = in.Contacts
		importedContacts  = make([]*mtproto.ImportedContact, 0, len(contacts))
		popularContactMap = make(map[string]*mtproto.TLPopularContact, len(contacts))
		updList           = make([]int64, 0, len(contacts))
		idList            = make([]int64, 0, len(contacts))
	)

	importContacts := make(map[string]*contactItem)
	// 1. Заполняем importContacts всеми импортированными контактами
	phoneList := make([]string, 0, len(contacts))
	for _, c2 := range contacts {
		phoneList = append(phoneList, c2.Phone)
		importContacts[c2.Phone] = &contactItem{unregistered: true, c: c2}
	}

	// 2. Ищем среди импортированных контактов те, которые зарегистрированны.
	registeredContacts, _ := c.svcCtx.Dao.UsersDAO.SelectUsersByPhoneList(c.ctx, phoneList)
	var contactIdList []int64

	// Непонятно, почему закомментированно
	// clear phoneList
	// phoneList = phoneList[0:0]
	for i := 0; i < len(registeredContacts); i++ {
		if c2, ok := importContacts[registeredContacts[i].Phone]; ok {
			c2.unregistered = false
			c2.userId = registeredContacts[i].Id
			// Зачем добавлять телефон в список, если он итак должен быть там?
			phoneList = append(phoneList, registeredContacts[i].Phone)
			contactIdList = append(contactIdList, registeredContacts[i].Id)
		} else {
			c2.unregistered = true
		}
	}

	if len(contactIdList) > 0 {
		// 3. Ищем среди импортированных и зарегистрированных контактов те,
		// которые уже являются моими контактами
		myContacts, _ := c.svcCtx.Dao.UserContactsDAO.SelectListByIdList(c.ctx, in.UserId, contactIdList)
		c.Logger.Infof("myContacts - %v", myContacts)
		for i := 0; i < len(myContacts); i++ {
			if c2, ok := importContacts[myContacts[i].ContactPhone]; ok {
				c2.contactId = myContacts[i].ContactUserId
			}
		}
	}

	if len(contactIdList) > 0 {
		// 4. Определяем, есть ли я в контактах у того, кого я добавляю
		importedMyContacts, _ := c.svcCtx.Dao.ImportedContactsDAO.SelectListByImportedList(c.ctx, in.UserId, contactIdList)
		c.Logger.Infof("importedMyContacts - %v", importedMyContacts)
		for i := 0; i < len(importedMyContacts); i++ {
			for _, c2 := range importContacts {
				if c2.userId == importedMyContacts[i].ImportedUserId {
					c2.importContactId = c2.userId
					break
				}
			}
		}
	}

	// clear phoneList
	phoneList = phoneList[0:0]
	for _, c2 := range importContacts {
		if c2.unregistered {
			go func() {
				// 1. Сохраняем куда-то приглашения незарегистрированных контактов
				// Что мы потом с ними будем делать?
				unregisteredContactsDO := &dataobject.UnregisteredContactsDO{
					Phone:           c2.c.Phone,
					ImporterUserId:  in.UserId,
					ImportFirstName: c2.c.FirstName,
					ImportLastName:  c2.c.LastName,
				}
				c.svcCtx.Dao.UnregisteredContactsDAO.InsertOrUpdate(c.ctx, unregisteredContactsDO)
			}()

			// Какая-то логика расчета популярных контактов. Очевидно, недоделанная
			//popularContactsDO := &dataobject.PopularContactsDO{
			//	Phone:     c2.c.Phone,
			//	Importers: 1,
			//}
			//c.dao.PopularContactsDAO.InsertOrUpdate(popularContactsDO)
			phoneList = append(phoneList, c2.c.Phone)
			popularContact := mtproto.MakeTLPopularContact(&mtproto.PopularContact{
				ClientId:  c2.c.ClientId,
				Importers: 1, // TODO(@benqi): get importers
			})
			popularContactMap[c2.c.Phone] = popularContact
			// &popularContactData{c2.c.Phone, c2.c.ClientId})
		} else {
			// Добавляем в контакты зарегистрированного пользователя
			userContactsDO := &dataobject.UserContactsDO{
				OwnerUserId:      in.UserId,
				ContactUserId:    c2.userId,
				ContactPhone:     c2.c.Phone,
				ContactFirstName: c2.c.FirstName,
				ContactLastName:  c2.c.LastName,
				Date2:            time.Now().Unix(),
			}

			if c2.contactId > 0 {
				if c2.importContactId > 0 {
					updList = append(updList, c2.importContactId)
				}

				// Если контакт уже существует, то обновляем для него
				// first_name, last_name
				c.svcCtx.Dao.UserContactsDAO.UpdateContactName(
					c.ctx,
					userContactsDO.ContactFirstName,
					userContactsDO.ContactLastName,
					userContactsDO.OwnerUserId,
					userContactsDO.ContactUserId)
			} else {
				userContactsDO.IsDeleted = false
				if c2.importContactId > 0 {
					// Существующий контакт стал взаимным. Обновляем его
					userContactsDO.Mutual = true

					updList = append(updList, c2.importContactId)

					c.svcCtx.Dao.UserContactsDAO.UpdateMutual(c.ctx, true, userContactsDO.ContactUserId, userContactsDO.OwnerUserId)
				} else {
					importedContactsDO := &dataobject.ImportedContactsDO{
						UserId:         userContactsDO.ContactUserId,
						ImportedUserId: userContactsDO.OwnerUserId,
					}
					c.svcCtx.Dao.ImportedContactsDAO.InsertOrUpdate(c.ctx, importedContactsDO)
				}
				c.svcCtx.Dao.UserContactsDAO.InsertOrUpdate(c.ctx, userContactsDO)
			}

			c.Logger.Infof("userContactsDO - %v", userContactsDO)
			c.Logger.Infof("c2 - %v", c2)

			importedContact := mtproto.MakeTLImportedContact(&mtproto.ImportedContact{
				UserId:   userContactsDO.ContactUserId,
				ClientId: c2.c.ClientId,
			})
			importedContacts = append(importedContacts, importedContact.To_ImportedContact())
			idList = append(idList, userContactsDO.ContactUserId)
		}
	}

	// Продолжение недоделанной логики про популярные контакты
	popularContacts := make([]*mtproto.PopularContact, 0, len(phoneList))
	if len(phoneList) > 0 {
		popularDOList, _ := c.svcCtx.Dao.PopularContactsDAO.SelectImportersList(c.ctx, phoneList)
		for i := 0; i < len(popularDOList); i++ {
			if c2, ok := popularContactMap[popularDOList[i].Phone]; ok {
				c2.SetImporters(popularDOList[i].Importers + 1)
			}
		}

		for _, c2 := range popularContactMap {
			popularContacts = append(popularContacts, c2.To_PopularContact())
		}

		go func() {
			// TODO:
			// m.PopularContactsDAO.IncreaseImportersList(context.Background(), phoneList)
		}()
	}

	users, _ := c.UserGetMutableUsers(&user.TLUserGetMutableUsers{
		Id: append(idList, in.UserId),
	})

	// importedContacts, popularContacts, updList
	rImportContacts := user.MakeTLUserImportedContacts(&user.UserImportedContacts{
		Imported:       importedContacts,
		PopularInvites: popularContacts,
		RetryContacts:  []int64{},
		Users:          users.GetUserListByIdList(in.UserId, idList...),
		UpdateIdList:   updList,
	}).To_UserImportedContacts()

	return rImportContacts, nil
}
