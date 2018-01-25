package services 

import (
    "context"
    "net/http"
    "database/sql"
    "github.com/Juniper/contrail/pkg/generated/models"
    "github.com/Juniper/contrail/pkg/generated/db"
    "github.com/satori/go.uuid"
    "github.com/labstack/echo"
    "github.com/Juniper/contrail/pkg/common"

	log "github.com/sirupsen/logrus"
)

//RESTCreateAnalyticsNode handle a Create REST service.
func (service *ContrailService) RESTCreateAnalyticsNode(c echo.Context) error {
    requestData := &models.AnalyticsNodeCreateRequest{
        AnalyticsNode: models.MakeAnalyticsNode(),
    }
    if err := c.Bind(requestData); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "analytics_node",
        }).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
    ctx := c.Request().Context()
    response, err := service.CreateAnalyticsNode(ctx, requestData)
    if err != nil {
        return common.ToHTTPError(err)
    } 
    return c.JSON(http.StatusCreated, response)
}

//CreateAnalyticsNode handle a Create API
func (service *ContrailService) CreateAnalyticsNode(
    ctx context.Context, 
    request *models.AnalyticsNodeCreateRequest) (*models.AnalyticsNodeCreateResponse, error) {
    model := request.AnalyticsNode
    if model.UUID == "" {
        model.UUID = uuid.NewV4().String()
    }

    if model.FQName == nil {
       return nil, common.ErrorBadRequest("Missing fq_name")
    }

    auth := common.GetAuthCTX(ctx)
    if auth == nil {
        return nil, common.ErrorUnauthenticated
    }
    model.Perms2.Owner = auth.ProjectID()
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.CreateAnalyticsNode(tx, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "analytics_node",
        }).Debug("db create failed on create")
       return nil, common.ErrorInternal 
    }
    return &models.AnalyticsNodeCreateResponse{
        AnalyticsNode: request.AnalyticsNode,
    }, nil
}

//RESTUpdateAnalyticsNode handles a REST Update request.
func (service *ContrailService) RESTUpdateAnalyticsNode(c echo.Context) error {
    id := c.Param("id")
    request := &models.AnalyticsNodeUpdateRequest{}
    if err := c.Bind(request); err != nil {
            log.WithFields(log.Fields{
                "err": err,
                "resource": "analytics_node",
            }).Debug("bind failed on update")
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
    }
    request.ID = id
    ctx := c.Request().Context()
    response, err := service.UpdateAnalyticsNode(ctx, request)
    if err != nil {
        return nil, common.ToHTTPError(err)
    }
    return c.JSON(http.StatusOK, response)
}

//UpdateAnalyticsNode handles a Update request.
func (service *ContrailService) UpdateAnalyticsNode(ctx context.Context, request *models.AnalyticsNodeUpdateRequest) (*models.AnalyticsNodeUpdateResponse, error) {
    id = request.ID
    model = request.AnalyticsNode
    if model == nil {
        return nil, common.ErrorBadRequest("Update body is empty")
    }
    auth := common.GetAuthCTX(ctx)
    ok := common.SetValueByPath(model, "Perms2.Owner", ".", auth.ProjectID())
    if !ok {
        return nil, common.ErrorBadRequest("Invalid JSON format")
    }
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.UpdateAnalyticsNode(tx, id, model)
        }); err != nil {
        log.WithFields(log.Fields{
            "err": err,
            "resource": "analytics_node",
        }).Debug("db update failed")
        return nil, common.ErrorInternal
    }
    return &models.AnalyticsNode.UpdateResponse{
        AnalyticsNode: model,
    }, nil
}

//RESTDeleteAnalyticsNode delete a resource using REST service.
func (service *ContrailService) RESTDeleteAnalyticsNode(c echo.Context) error {
    id := c.Param("id")
    request := &models.AnalyticsNodeDeleteRequest{
        ID: id
    } 
    ctx := c.Request().Context()
    response, err := service.DeleteAnalyticsNode(ctx, request)
    if err != nil {
        return common.ToHTTPError(err)
    }
    return c.JSON(http.StatusNoContent, nil)
}

//DeleteAnalyticsNode delete a resource.
func (service *ContrailService) DeleteAnalyticsNode(ctx context.Context, request *models.AnalyticsNodeDeleteRequest) (*models.AnalyticsNodeDeleteResponse, error) {
    id := request.ID
    auth := common.GetAuthCTX(ctx)
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            return db.DeleteAnalyticsNode(tx, id, auth)
        }); err != nil {
            log.WithField("err", err).Debug("error deleting a resource")
        return nil, common.ErrorInternal
    }
    return &models.AnalyticsNodeDeleteResponse{
        ID: id,
    }, nil
}

//RESTShowAnalyticsNode a REST Show request.
func (service *ContrailService) RESTShowAnalyticsNode(c echo.Context) (error) {
    id := c.Param("id")
    auth := common.GetAuthContext(c)
    var result []*models.AnalyticsNode
    var err error
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListAnalyticsNode(tx, &common.ListSpec{
                Limit: 1,
                Auth: auth,
                Filter: common.Filter{
                    "uuid": []string{id},
                },
            })
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "analytics_node": result,
    })
}

//RESTListAnalyticsNode handles a List REST service Request.
func (service *ContrailService) RESTListAnalyticsNode(c echo.Context) (error) {
    var result []*models.AnalyticsNode
    var err error
    auth := common.GetAuthContext(c)
    listSpec := common.GetListSpec(c)
    listSpec.Auth = auth
    if err := common.DoInTransaction(
        service.DB,
        func (tx *sql.Tx) error {
            result, err = db.ListAnalyticsNode(tx, listSpec)
            return err
        }); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    return c.JSON(http.StatusOK, map[string]interface{}{
        "analytics-nodes": result,
    })
}