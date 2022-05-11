// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package core

import (
	"regexp"
)

import (
	"github.com/teamgram/proto/mtproto"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

// ContactsImportContacts
// contacts.importContacts#2c800be5 contacts:Vector<InputContact> = contacts.ImportedContacts;
func (c *ContactsCore) ContactsImportContacts(in *mtproto.TLContactsImportContacts) (*mtproto.Contacts_ImportedContacts, error) {
	c.Logger.Infof("contacts.importContacts: %#v", in)

	l := len(in.Contacts)
	if l > 1 {
		// Будет обработан только один запрос
		c.Logger.Errorf("contacts.ImportContacts - warning: len is %d", l)
	}
		
	for _, inContact := range in.Contacts {
		// 400	CONTACT_NAME_EMPTY	Contact name empty.
		if inContact.FirstName == "" && inContact.LastName == "" {
			err := mtproto.ErrContactNameEmpty
			c.Logger.Errorf("contacts.importContacts - empty names error: %v", err)
			return nil, err
		}
	}
	for _, c := range in.Contacts {
		reg, _ := regexp.Compile("[^0-9]+")
		c.Phone = reg.ReplaceAllString(c.Phone, "")
	}
	contacts, err := c.svcCtx.Dao.UserClient.UserImportContacts(
		c.ctx, &userpb.TLUserImportContacts{
			UserId: c.MD.UserId,
			Contacts: in.Contacts})
	
	if err != nil {
		c.Logger.Errorf("contacts.importContacts - error: %v", err)
		return nil, err
	}

	return mtproto.MakeTLContactsImportedContacts(&mtproto.Contacts_ImportedContacts{
		Imported:       contacts.Imported,
		PopularInvites: contacts.PopularInvites,
		RetryContacts:  contacts.RetryContacts,
		Users:          contacts.Users,
	}).To_Contacts_ImportedContacts(), nil
}


