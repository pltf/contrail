package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertAnalyticsNodeQuery = "insert into `analytics_node` (`analytics_node_ip_address`,`display_name`,`key_value_pair`,`share`,`owner`,`owner_access`,`global_access`,`uuid`,`fq_name`,`enable`,`description`,`created`,`creator`,`user_visible`,`last_modified`,`group_access`,`permissions_owner`,`permissions_owner_access`,`other_access`,`group`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateAnalyticsNodeQuery = "update `analytics_node` set `analytics_node_ip_address` = ?,`display_name` = ?,`key_value_pair` = ?,`share` = ?,`owner` = ?,`owner_access` = ?,`global_access` = ?,`uuid` = ?,`fq_name` = ?,`enable` = ?,`description` = ?,`created` = ?,`creator` = ?,`user_visible` = ?,`last_modified` = ?,`group_access` = ?,`permissions_owner` = ?,`permissions_owner_access` = ?,`other_access` = ?,`group` = ?;"
const deleteAnalyticsNodeQuery = "delete from `analytics_node` where uuid = ?"

// AnalyticsNodeFields is db columns for AnalyticsNode
var AnalyticsNodeFields = []string{
	"analytics_node_ip_address",
	"display_name",
	"key_value_pair",
	"share",
	"owner",
	"owner_access",
	"global_access",
	"uuid",
	"fq_name",
	"enable",
	"description",
	"created",
	"creator",
	"user_visible",
	"last_modified",
	"group_access",
	"permissions_owner",
	"permissions_owner_access",
	"other_access",
	"group",
}

// AnalyticsNodeRefFields is db reference fields for AnalyticsNode
var AnalyticsNodeRefFields = map[string][]string{}

// CreateAnalyticsNode inserts AnalyticsNode to DB
func CreateAnalyticsNode(tx *sql.Tx, model *models.AnalyticsNode) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertAnalyticsNodeQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertAnalyticsNodeQuery,
	}).Debug("create query")
	_, err = stmt.Exec(string(model.AnalyticsNodeIPAddress),
		string(model.DisplayName),
		common.MustJSON(model.Annotations.KeyValuePair),
		common.MustJSON(model.Perms2.Share),
		string(model.Perms2.Owner),
		int(model.Perms2.OwnerAccess),
		int(model.Perms2.GlobalAccess),
		string(model.UUID),
		common.MustJSON(model.FQName),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Created),
		string(model.IDPerms.Creator),
		bool(model.IDPerms.UserVisible),
		string(model.IDPerms.LastModified),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OwnerAccess),
		int(model.IDPerms.Permissions.OtherAccess),
		string(model.IDPerms.Permissions.Group))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanAnalyticsNode(values map[string]interface{}) (*models.AnalyticsNode, error) {
	m := models.MakeAnalyticsNode()

	if value, ok := values["analytics_node_ip_address"]; ok {

		castedValue := common.InterfaceToString(value)

		m.AnalyticsNodeIPAddress = models.IpAddressType(castedValue)

	}

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["global_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["description"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Description = castedValue

	}

	if value, ok := values["created"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Created = castedValue

	}

	if value, ok := values["creator"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Creator = castedValue

	}

	if value, ok := values["user_visible"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.UserVisible = castedValue

	}

	if value, ok := values["last_modified"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

	}

	if value, ok := values["group_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

	}

	if value, ok := values["permissions_owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["permissions_owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["other_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Group = castedValue

	}

	return m, nil
}

// ListAnalyticsNode lists AnalyticsNode with list spec.
func ListAnalyticsNode(tx *sql.Tx, spec *common.ListSpec) ([]*models.AnalyticsNode, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "analytics_node"
	spec.Fields = AnalyticsNodeFields
	spec.RefFields = AnalyticsNodeRefFields
	result := models.MakeAnalyticsNodeSlice()
	query, columns, values := common.BuildListQuery(spec)
	log.WithFields(log.Fields{
		"listSpec": spec,
		"query":    query,
	}).Debug("select query")
	rows, err = tx.Query(query, values...)
	if err != nil {
		return nil, errors.Wrap(err, "select query failed")
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "row error")
	}
	for rows.Next() {
		valuesMap := map[string]interface{}{}
		values := make([]interface{}, len(columns))
		valuesPointers := make([]interface{}, len(columns))
		for _, index := range columns {
			valuesPointers[index] = &values[index]
		}
		if err := rows.Scan(valuesPointers...); err != nil {
			return nil, errors.Wrap(err, "scan failed")
		}
		for column, index := range columns {
			val := valuesPointers[index].(*interface{})
			valuesMap[column] = *val
		}
		log.WithFields(log.Fields{
			"valuesMap": valuesMap,
		}).Debug("valueMap")
		m, err := scanAnalyticsNode(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// ShowAnalyticsNode shows AnalyticsNode resource
func ShowAnalyticsNode(tx *sql.Tx, uuid string) (*models.AnalyticsNode, error) {
	list, err := ListAnalyticsNode(tx, &common.ListSpec{
		Filter: map[string]interface{}{"uuid": uuid},
		Limit:  1})
	if len(list) == 0 {
		return nil, errors.Wrap(err, "show query failed")
	}
	return list[0], err
}

// UpdateAnalyticsNode updates a resource
func UpdateAnalyticsNode(tx *sql.Tx, uuid string, model *models.AnalyticsNode) error {
	//TODO(nati) support update
	return nil
}

// DeleteAnalyticsNode deletes a resource
func DeleteAnalyticsNode(tx *sql.Tx, uuid string) error {
	stmt, err := tx.Prepare(deleteAnalyticsNodeQuery)
	if err != nil {
		return errors.Wrap(err, "preparing delete query failed")
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	if err != nil {
		return errors.Wrap(err, "delete failed")
	}
	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("deleted")
	return nil
}
