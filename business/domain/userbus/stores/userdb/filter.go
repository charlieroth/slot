package userdb

import (
	"bytes"
	"strings"

	"github.com/charlieroth/slot/business/domain/userbus"
)

func (s *Store) applyFilter(filter userbus.QueryFilter, data map[string]any, buf *bytes.Buffer) {
	var wc []string

	if filter.ID != nil {
		wc = append(wc, "id = @id")
		data["id"] = *filter.ID
	}

	if filter.Email != nil {
		wc = append(wc, "email = @email")
		data["email"] = *filter.Email
	}

	if filter.Name != nil {
		wc = append(wc, "name = @name")
		data["name"] = *filter.Name
	}

	if filter.Phone != nil {
		wc = append(wc, "phone = @phone")
		data["phone"] = *filter.Phone
	}

	if filter.UserType != nil {
		wc = append(wc, "user_type = @user_type")
		data["user_type"] = *filter.UserType
	}

	if filter.TimeZone != nil {
		wc = append(wc, "time_zone = @time_zone")
		data["time_zone"] = *filter.TimeZone
	}

	if len(wc) > 0 {
		buf.WriteString("WHERE ")
		buf.WriteString(strings.Join(wc, " AND "))
	}
}
