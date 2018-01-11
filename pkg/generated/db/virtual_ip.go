package db

import (
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertVirtualIPQuery = "insert into `virtual_ip` (`subnet_id`,`status_description`,`status`,`protocol_port`,`protocol`,`persistence_type`,`persistence_cookie_name`,`connection_limit`,`admin_state`,`address`,`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const updateVirtualIPQuery = "update `virtual_ip` set `subnet_id` = ?,`status_description` = ?,`status` = ?,`protocol_port` = ?,`protocol` = ?,`persistence_type` = ?,`persistence_cookie_name` = ?,`connection_limit` = ?,`admin_state` = ?,`address` = ?,`uuid` = ?,`share` = ?,`owner_access` = ?,`owner` = ?,`global_access` = ?,`parent_uuid` = ?,`parent_type` = ?,`user_visible` = ?,`permissions_owner_access` = ?,`permissions_owner` = ?,`other_access` = ?,`group_access` = ?,`group` = ?,`last_modified` = ?,`enable` = ?,`description` = ?,`creator` = ?,`created` = ?,`fq_name` = ?,`display_name` = ?,`key_value_pair` = ?;"
const deleteVirtualIPQuery = "delete from `virtual_ip` where uuid = ?"

// VirtualIPFields is db columns for VirtualIP
var VirtualIPFields = []string{
	"subnet_id",
	"status_description",
	"status",
	"protocol_port",
	"protocol",
	"persistence_type",
	"persistence_cookie_name",
	"connection_limit",
	"admin_state",
	"address",
	"uuid",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"user_visible",
	"permissions_owner_access",
	"permissions_owner",
	"other_access",
	"group_access",
	"group",
	"last_modified",
	"enable",
	"description",
	"creator",
	"created",
	"fq_name",
	"display_name",
	"key_value_pair",
}

// VirtualIPRefFields is db reference fields for VirtualIP
var VirtualIPRefFields = map[string][]string{

	"loadbalancer_pool": {
	// <common.Schema Value>

	},

	"virtual_machine_interface": {
	// <common.Schema Value>

	},
}

// VirtualIPBackRefFields is db back reference fields for VirtualIP
var VirtualIPBackRefFields = map[string][]string{}

// VirtualIPParentTypes is possible parents for VirtualIP
var VirtualIPParents = []string{

	"project",
}

const insertVirtualIPLoadbalancerPoolQuery = "insert into `ref_virtual_ip_loadbalancer_pool` (`from`, `to` ) values (?, ?);"

const insertVirtualIPVirtualMachineInterfaceQuery = "insert into `ref_virtual_ip_virtual_machine_interface` (`from`, `to` ) values (?, ?);"

// CreateVirtualIP inserts VirtualIP to DB
func CreateVirtualIP(tx *sql.Tx, model *models.VirtualIP) error {
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertVirtualIPQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertVirtualIPQuery,
	}).Debug("create query")
	_, err = stmt.Exec(string(model.VirtualIPProperties.SubnetID),
		string(model.VirtualIPProperties.StatusDescription),
		string(model.VirtualIPProperties.Status),
		int(model.VirtualIPProperties.ProtocolPort),
		string(model.VirtualIPProperties.Protocol),
		string(model.VirtualIPProperties.PersistenceType),
		string(model.VirtualIPProperties.PersistenceCookieName),
		int(model.VirtualIPProperties.ConnectionLimit),
		bool(model.VirtualIPProperties.AdminState),
		string(model.VirtualIPProperties.Address),
		string(model.UUID),
		common.MustJSON(model.Perms2.Share),
		int(model.Perms2.OwnerAccess),
		string(model.Perms2.Owner),
		int(model.Perms2.GlobalAccess),
		string(model.ParentUUID),
		string(model.ParentType),
		bool(model.IDPerms.UserVisible),
		int(model.IDPerms.Permissions.OwnerAccess),
		string(model.IDPerms.Permissions.Owner),
		int(model.IDPerms.Permissions.OtherAccess),
		int(model.IDPerms.Permissions.GroupAccess),
		string(model.IDPerms.Permissions.Group),
		string(model.IDPerms.LastModified),
		bool(model.IDPerms.Enable),
		string(model.IDPerms.Description),
		string(model.IDPerms.Creator),
		string(model.IDPerms.Created),
		common.MustJSON(model.FQName),
		string(model.DisplayName),
		common.MustJSON(model.Annotations.KeyValuePair))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	stmtLoadbalancerPoolRef, err := tx.Prepare(insertVirtualIPLoadbalancerPoolQuery)
	if err != nil {
		return errors.Wrap(err, "preparing LoadbalancerPoolRefs create statement failed")
	}
	defer stmtLoadbalancerPoolRef.Close()
	for _, ref := range model.LoadbalancerPoolRefs {

		_, err = stmtLoadbalancerPoolRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "LoadbalancerPoolRefs create failed")
		}
	}

	stmtVirtualMachineInterfaceRef, err := tx.Prepare(insertVirtualIPVirtualMachineInterfaceQuery)
	if err != nil {
		return errors.Wrap(err, "preparing VirtualMachineInterfaceRefs create statement failed")
	}
	defer stmtVirtualMachineInterfaceRef.Close()
	for _, ref := range model.VirtualMachineInterfaceRefs {

		_, err = stmtVirtualMachineInterfaceRef.Exec(model.UUID, ref.UUID)
		if err != nil {
			return errors.Wrap(err, "VirtualMachineInterfaceRefs create failed")
		}
	}

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "virtual_ip",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return err
}

func scanVirtualIP(values map[string]interface{}) (*models.VirtualIP, error) {
	m := models.MakeVirtualIP()

	if value, ok := values["subnet_id"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VirtualIPProperties.SubnetID = models.UuidStringType(castedValue)

	}

	if value, ok := values["status_description"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VirtualIPProperties.StatusDescription = castedValue

	}

	if value, ok := values["status"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VirtualIPProperties.Status = castedValue

	}

	if value, ok := values["protocol_port"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.VirtualIPProperties.ProtocolPort = castedValue

	}

	if value, ok := values["protocol"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VirtualIPProperties.Protocol = models.LoadbalancerProtocolType(castedValue)

	}

	if value, ok := values["persistence_type"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VirtualIPProperties.PersistenceType = models.SessionPersistenceType(castedValue)

	}

	if value, ok := values["persistence_cookie_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VirtualIPProperties.PersistenceCookieName = castedValue

	}

	if value, ok := values["connection_limit"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.VirtualIPProperties.ConnectionLimit = castedValue

	}

	if value, ok := values["admin_state"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.VirtualIPProperties.AdminState = castedValue

	}

	if value, ok := values["address"]; ok {

		castedValue := common.InterfaceToString(value)

		m.VirtualIPProperties.Address = models.IpAddressType(castedValue)

	}

	if value, ok := values["uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.UUID = castedValue

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.Perms2.Owner = castedValue

	}

	if value, ok := values["global_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.Perms2.GlobalAccess = models.AccessType(castedValue)

	}

	if value, ok := values["parent_uuid"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ParentUUID = castedValue

	}

	if value, ok := values["parent_type"]; ok {

		castedValue := common.InterfaceToString(value)

		m.ParentType = castedValue

	}

	if value, ok := values["user_visible"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.UserVisible = castedValue

	}

	if value, ok := values["permissions_owner_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OwnerAccess = models.AccessType(castedValue)

	}

	if value, ok := values["permissions_owner"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Owner = castedValue

	}

	if value, ok := values["other_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.OtherAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group_access"]; ok {

		castedValue := common.InterfaceToInt(value)

		m.IDPerms.Permissions.GroupAccess = models.AccessType(castedValue)

	}

	if value, ok := values["group"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Permissions.Group = castedValue

	}

	if value, ok := values["last_modified"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.LastModified = castedValue

	}

	if value, ok := values["enable"]; ok {

		castedValue := common.InterfaceToBool(value)

		m.IDPerms.Enable = castedValue

	}

	if value, ok := values["description"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Description = castedValue

	}

	if value, ok := values["creator"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Creator = castedValue

	}

	if value, ok := values["created"]; ok {

		castedValue := common.InterfaceToString(value)

		m.IDPerms.Created = castedValue

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["display_name"]; ok {

		castedValue := common.InterfaceToString(value)

		m.DisplayName = castedValue

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	if value, ok := values["ref_loadbalancer_pool"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualIPLoadbalancerPoolRef{}
			referenceModel.UUID = uuid
			m.LoadbalancerPoolRefs = append(m.LoadbalancerPoolRefs, referenceModel)

		}
	}

	if value, ok := values["ref_virtual_machine_interface"]; ok {
		var references []interface{}
		stringValue := common.InterfaceToString(value)
		json.Unmarshal([]byte("["+stringValue+"]"), &references)
		for _, reference := range references {
			referenceMap, ok := reference.(map[string]interface{})
			if !ok {
				continue
			}
			uuid := common.InterfaceToString(referenceMap["to"])
			if uuid == "" {
				continue
			}
			referenceModel := &models.VirtualIPVirtualMachineInterfaceRef{}
			referenceModel.UUID = uuid
			m.VirtualMachineInterfaceRefs = append(m.VirtualMachineInterfaceRefs, referenceModel)

		}
	}

	return m, nil
}

// ListVirtualIP lists VirtualIP with list spec.
func ListVirtualIP(tx *sql.Tx, spec *common.ListSpec) ([]*models.VirtualIP, error) {
	var rows *sql.Rows
	var err error
	//TODO (check input)
	spec.Table = "virtual_ip"
	spec.Fields = VirtualIPFields
	spec.RefFields = VirtualIPRefFields
	spec.BackRefFields = VirtualIPBackRefFields
	result := models.MakeVirtualIPSlice()

	if spec.ParentFQName != nil {
		parentMetaData, err := common.GetMetaData(tx, "", spec.ParentFQName)
		if err != nil {
			return nil, errors.Wrap(err, "can't find parents")
		}
		spec.Filter.AppendValues("parent_uuid", []string{parentMetaData.UUID})
	}

	query := spec.BuildQuery()
	columns := spec.Columns
	values := spec.Values
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
		m, err := scanVirtualIP(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	return result, nil
}

// UpdateVirtualIP updates a resource
func UpdateVirtualIP(tx *sql.Tx, uuid string, model *models.VirtualIP) error {
	//TODO(nati) support update
	return nil
}

// DeleteVirtualIP deletes a resource
func DeleteVirtualIP(tx *sql.Tx, uuid string, auth *common.AuthContext) error {
	query := deleteVirtualIPQuery
	var err error

	if auth.IsAdmin() {
		_, err = tx.Exec(query, uuid)
	} else {
		query += " and owner = ?"
		_, err = tx.Exec(query, uuid, auth.ProjectID())
	}

	if err != nil {
		return errors.Wrap(err, "delete failed")
	}

	err = common.DeleteMetaData(tx, uuid)
	log.WithFields(log.Fields{
		"uuid": uuid,
	}).Debug("deleted")
	return err
}
