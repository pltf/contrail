package db

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/generated/models"
	"github.com/Juniper/contrail/pkg/schema"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const insertLoadbalancerHealthmonitorQuery = "insert into `loadbalancer_healthmonitor` (`uuid`,`share`,`owner_access`,`owner`,`global_access`,`parent_uuid`,`parent_type`,`url_path`,`timeout`,`monitor_type`,`max_retries`,`http_method`,`expected_codes`,`delay`,`admin_state`,`user_visible`,`permissions_owner_access`,`permissions_owner`,`other_access`,`group_access`,`group`,`last_modified`,`enable`,`description`,`creator`,`created`,`fq_name`,`display_name`,`key_value_pair`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
const deleteLoadbalancerHealthmonitorQuery = "delete from `loadbalancer_healthmonitor` where uuid = ?"

// LoadbalancerHealthmonitorFields is db columns for LoadbalancerHealthmonitor
var LoadbalancerHealthmonitorFields = []string{
	"uuid",
	"share",
	"owner_access",
	"owner",
	"global_access",
	"parent_uuid",
	"parent_type",
	"url_path",
	"timeout",
	"monitor_type",
	"max_retries",
	"http_method",
	"expected_codes",
	"delay",
	"admin_state",
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

// LoadbalancerHealthmonitorRefFields is db reference fields for LoadbalancerHealthmonitor
var LoadbalancerHealthmonitorRefFields = map[string][]string{}

// LoadbalancerHealthmonitorBackRefFields is db back reference fields for LoadbalancerHealthmonitor
var LoadbalancerHealthmonitorBackRefFields = map[string][]string{}

// LoadbalancerHealthmonitorParentTypes is possible parents for LoadbalancerHealthmonitor
var LoadbalancerHealthmonitorParents = []string{

	"project",
}

// CreateLoadbalancerHealthmonitor inserts LoadbalancerHealthmonitor to DB
func (db *DB) createLoadbalancerHealthmonitor(
	ctx context.Context,
	request *models.CreateLoadbalancerHealthmonitorRequest) error {
	tx := common.GetTransaction(ctx)
	model := request.LoadbalancerHealthmonitor
	// Prepare statement for inserting data
	stmt, err := tx.Prepare(insertLoadbalancerHealthmonitorQuery)
	if err != nil {
		return errors.Wrap(err, "preparing create statement failed")
	}
	defer stmt.Close()
	log.WithFields(log.Fields{
		"model": model,
		"query": insertLoadbalancerHealthmonitorQuery,
	}).Debug("create query")
	_, err = stmt.ExecContext(ctx, string(model.GetUUID()),
		common.MustJSON(model.GetPerms2().GetShare()),
		int(model.GetPerms2().GetOwnerAccess()),
		string(model.GetPerms2().GetOwner()),
		int(model.GetPerms2().GetGlobalAccess()),
		string(model.GetParentUUID()),
		string(model.GetParentType()),
		string(model.GetLoadbalancerHealthmonitorProperties().GetURLPath()),
		int(model.GetLoadbalancerHealthmonitorProperties().GetTimeout()),
		string(model.GetLoadbalancerHealthmonitorProperties().GetMonitorType()),
		int(model.GetLoadbalancerHealthmonitorProperties().GetMaxRetries()),
		string(model.GetLoadbalancerHealthmonitorProperties().GetHTTPMethod()),
		string(model.GetLoadbalancerHealthmonitorProperties().GetExpectedCodes()),
		int(model.GetLoadbalancerHealthmonitorProperties().GetDelay()),
		bool(model.GetLoadbalancerHealthmonitorProperties().GetAdminState()),
		bool(model.GetIDPerms().GetUserVisible()),
		int(model.GetIDPerms().GetPermissions().GetOwnerAccess()),
		string(model.GetIDPerms().GetPermissions().GetOwner()),
		int(model.GetIDPerms().GetPermissions().GetOtherAccess()),
		int(model.GetIDPerms().GetPermissions().GetGroupAccess()),
		string(model.GetIDPerms().GetPermissions().GetGroup()),
		string(model.GetIDPerms().GetLastModified()),
		bool(model.GetIDPerms().GetEnable()),
		string(model.GetIDPerms().GetDescription()),
		string(model.GetIDPerms().GetCreator()),
		string(model.GetIDPerms().GetCreated()),
		common.MustJSON(model.GetFQName()),
		string(model.GetDisplayName()),
		common.MustJSON(model.GetAnnotations().GetKeyValuePair()))
	if err != nil {
		return errors.Wrap(err, "create failed")
	}

	metaData := &common.MetaData{
		UUID:   model.UUID,
		Type:   "loadbalancer_healthmonitor",
		FQName: model.FQName,
	}
	err = common.CreateMetaData(tx, metaData)
	if err != nil {
		return err
	}
	err = common.CreateSharing(tx, "loadbalancer_healthmonitor", model.UUID, model.GetPerms2().GetShare())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"model": model,
	}).Debug("created")
	return nil
}

func scanLoadbalancerHealthmonitor(values map[string]interface{}) (*models.LoadbalancerHealthmonitor, error) {
	m := models.MakeLoadbalancerHealthmonitor()

	if value, ok := values["uuid"]; ok {

		m.UUID = schema.InterfaceToString(value)

	}

	if value, ok := values["share"]; ok {

		json.Unmarshal(value.([]byte), &m.Perms2.Share)

	}

	if value, ok := values["owner_access"]; ok {

		m.Perms2.OwnerAccess = schema.InterfaceToInt64(value)

	}

	if value, ok := values["owner"]; ok {

		m.Perms2.Owner = schema.InterfaceToString(value)

	}

	if value, ok := values["global_access"]; ok {

		m.Perms2.GlobalAccess = schema.InterfaceToInt64(value)

	}

	if value, ok := values["parent_uuid"]; ok {

		m.ParentUUID = schema.InterfaceToString(value)

	}

	if value, ok := values["parent_type"]; ok {

		m.ParentType = schema.InterfaceToString(value)

	}

	if value, ok := values["url_path"]; ok {

		m.LoadbalancerHealthmonitorProperties.URLPath = schema.InterfaceToString(value)

	}

	if value, ok := values["timeout"]; ok {

		m.LoadbalancerHealthmonitorProperties.Timeout = schema.InterfaceToInt64(value)

	}

	if value, ok := values["monitor_type"]; ok {

		m.LoadbalancerHealthmonitorProperties.MonitorType = schema.InterfaceToString(value)

	}

	if value, ok := values["max_retries"]; ok {

		m.LoadbalancerHealthmonitorProperties.MaxRetries = schema.InterfaceToInt64(value)

	}

	if value, ok := values["http_method"]; ok {

		m.LoadbalancerHealthmonitorProperties.HTTPMethod = schema.InterfaceToString(value)

	}

	if value, ok := values["expected_codes"]; ok {

		m.LoadbalancerHealthmonitorProperties.ExpectedCodes = schema.InterfaceToString(value)

	}

	if value, ok := values["delay"]; ok {

		m.LoadbalancerHealthmonitorProperties.Delay = schema.InterfaceToInt64(value)

	}

	if value, ok := values["admin_state"]; ok {

		m.LoadbalancerHealthmonitorProperties.AdminState = schema.InterfaceToBool(value)

	}

	if value, ok := values["user_visible"]; ok {

		m.IDPerms.UserVisible = schema.InterfaceToBool(value)

	}

	if value, ok := values["permissions_owner_access"]; ok {

		m.IDPerms.Permissions.OwnerAccess = schema.InterfaceToInt64(value)

	}

	if value, ok := values["permissions_owner"]; ok {

		m.IDPerms.Permissions.Owner = schema.InterfaceToString(value)

	}

	if value, ok := values["other_access"]; ok {

		m.IDPerms.Permissions.OtherAccess = schema.InterfaceToInt64(value)

	}

	if value, ok := values["group_access"]; ok {

		m.IDPerms.Permissions.GroupAccess = schema.InterfaceToInt64(value)

	}

	if value, ok := values["group"]; ok {

		m.IDPerms.Permissions.Group = schema.InterfaceToString(value)

	}

	if value, ok := values["last_modified"]; ok {

		m.IDPerms.LastModified = schema.InterfaceToString(value)

	}

	if value, ok := values["enable"]; ok {

		m.IDPerms.Enable = schema.InterfaceToBool(value)

	}

	if value, ok := values["description"]; ok {

		m.IDPerms.Description = schema.InterfaceToString(value)

	}

	if value, ok := values["creator"]; ok {

		m.IDPerms.Creator = schema.InterfaceToString(value)

	}

	if value, ok := values["created"]; ok {

		m.IDPerms.Created = schema.InterfaceToString(value)

	}

	if value, ok := values["fq_name"]; ok {

		json.Unmarshal(value.([]byte), &m.FQName)

	}

	if value, ok := values["display_name"]; ok {

		m.DisplayName = schema.InterfaceToString(value)

	}

	if value, ok := values["key_value_pair"]; ok {

		json.Unmarshal(value.([]byte), &m.Annotations.KeyValuePair)

	}

	return m, nil
}

// ListLoadbalancerHealthmonitor lists LoadbalancerHealthmonitor with list spec.
func (db *DB) listLoadbalancerHealthmonitor(ctx context.Context, request *models.ListLoadbalancerHealthmonitorRequest) (response *models.ListLoadbalancerHealthmonitorResponse, err error) {
	var rows *sql.Rows
	tx := common.GetTransaction(ctx)
	qb := &common.ListQueryBuilder{}
	qb.Auth = common.GetAuthCTX(ctx)
	spec := request.Spec
	qb.Spec = spec
	qb.Table = "loadbalancer_healthmonitor"
	qb.Fields = LoadbalancerHealthmonitorFields
	qb.RefFields = LoadbalancerHealthmonitorRefFields
	qb.BackRefFields = LoadbalancerHealthmonitorBackRefFields
	result := []*models.LoadbalancerHealthmonitor{}

	if spec.ParentFQName != nil {
		parentMetaData, err := common.GetMetaData(tx, "", spec.ParentFQName)
		if err != nil {
			return nil, errors.Wrap(err, "can't find parents")
		}
		spec.Filters = common.AppendFilter(spec.Filters, "parent_uuid", parentMetaData.UUID)
	}

	query := qb.BuildQuery()
	columns := qb.Columns
	values := qb.Values
	log.WithFields(log.Fields{
		"listSpec": spec,
		"query":    query,
	}).Debug("select query")
	rows, err = tx.QueryContext(ctx, query, values...)
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
		m, err := scanLoadbalancerHealthmonitor(valuesMap)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}
		result = append(result, m)
	}
	response = &models.ListLoadbalancerHealthmonitorResponse{
		LoadbalancerHealthmonitors: result,
	}
	return response, nil
}

// UpdateLoadbalancerHealthmonitor updates a resource
func (db *DB) updateLoadbalancerHealthmonitor(
	ctx context.Context,
	request *models.UpdateLoadbalancerHealthmonitorRequest,
) error {
	//TODO
	return nil
}

// DeleteLoadbalancerHealthmonitor deletes a resource
func (db *DB) deleteLoadbalancerHealthmonitor(
	ctx context.Context,
	request *models.DeleteLoadbalancerHealthmonitorRequest) error {
	deleteQuery := deleteLoadbalancerHealthmonitorQuery
	selectQuery := "select count(uuid) from loadbalancer_healthmonitor where uuid = ?"
	var err error
	var count int
	uuid := request.ID
	tx := common.GetTransaction(ctx)
	auth := common.GetAuthCTX(ctx)
	if auth.IsAdmin() {
		row := tx.QueryRowContext(ctx, selectQuery, uuid)
		if err != nil {
			return errors.Wrap(err, "not found")
		}
		row.Scan(&count)
		if count == 0 {
			return errors.New("Not found")
		}
		_, err = tx.ExecContext(ctx, deleteQuery, uuid)
	} else {
		deleteQuery += " and owner = ?"
		selectQuery += " and owner = ?"
		row := tx.QueryRowContext(ctx, selectQuery, uuid, auth.ProjectID())
		if err != nil {
			return errors.Wrap(err, "not found")
		}
		row.Scan(&count)
		if count == 0 {
			return errors.New("Not found")
		}
		_, err = tx.ExecContext(ctx, deleteQuery, uuid, auth.ProjectID())
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

//CreateLoadbalancerHealthmonitor handle a Create API
func (db *DB) CreateLoadbalancerHealthmonitor(
	ctx context.Context,
	request *models.CreateLoadbalancerHealthmonitorRequest) (*models.CreateLoadbalancerHealthmonitorResponse, error) {
	model := request.LoadbalancerHealthmonitor
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.createLoadbalancerHealthmonitor(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer_healthmonitor",
		}).Debug("db create failed on create")
		return nil, common.ErrorInternal
	}
	return &models.CreateLoadbalancerHealthmonitorResponse{
		LoadbalancerHealthmonitor: request.LoadbalancerHealthmonitor,
	}, nil
}

//UpdateLoadbalancerHealthmonitor handles a Update request.
func (db *DB) UpdateLoadbalancerHealthmonitor(
	ctx context.Context,
	request *models.UpdateLoadbalancerHealthmonitorRequest) (*models.UpdateLoadbalancerHealthmonitorResponse, error) {
	model := request.LoadbalancerHealthmonitor
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.updateLoadbalancerHealthmonitor(ctx, request)
		}); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "loadbalancer_healthmonitor",
		}).Debug("db update failed")
		return nil, common.ErrorInternal
	}
	return &models.UpdateLoadbalancerHealthmonitorResponse{
		LoadbalancerHealthmonitor: model,
	}, nil
}

//DeleteLoadbalancerHealthmonitor delete a resource.
func (db *DB) DeleteLoadbalancerHealthmonitor(ctx context.Context, request *models.DeleteLoadbalancerHealthmonitorRequest) (*models.DeleteLoadbalancerHealthmonitorResponse, error) {
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			return db.deleteLoadbalancerHealthmonitor(ctx, request)
		}); err != nil {
		log.WithField("err", err).Debug("error deleting a resource")
		return nil, common.ErrorInternal
	}
	return &models.DeleteLoadbalancerHealthmonitorResponse{
		ID: request.ID,
	}, nil
}

//GetLoadbalancerHealthmonitor a Get request.
func (db *DB) GetLoadbalancerHealthmonitor(ctx context.Context, request *models.GetLoadbalancerHealthmonitorRequest) (response *models.GetLoadbalancerHealthmonitorResponse, err error) {
	spec := &models.ListSpec{
		Limit:  1,
		Detail: true,
		Filters: []*models.Filter{
			&models.Filter{
				Key:    "uuid",
				Values: []string{request.ID},
			},
		},
	}
	listRequest := &models.ListLoadbalancerHealthmonitorRequest{
		Spec: spec,
	}
	var result *models.ListLoadbalancerHealthmonitorResponse
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			result, err = db.listLoadbalancerHealthmonitor(ctx, listRequest)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	if len(result.LoadbalancerHealthmonitors) == 0 {
		return nil, common.ErrorNotFound
	}
	response = &models.GetLoadbalancerHealthmonitorResponse{
		LoadbalancerHealthmonitor: result.LoadbalancerHealthmonitors[0],
	}
	return response, nil
}

//ListLoadbalancerHealthmonitor handles a List service Request.
func (db *DB) ListLoadbalancerHealthmonitor(
	ctx context.Context,
	request *models.ListLoadbalancerHealthmonitorRequest) (response *models.ListLoadbalancerHealthmonitorResponse, err error) {
	if err := common.DoInTransaction(
		ctx,
		db.DB,
		func(ctx context.Context) error {
			response, err = db.listLoadbalancerHealthmonitor(ctx, request)
			return err
		}); err != nil {
		return nil, common.ErrorInternal
	}
	return response, nil
}
