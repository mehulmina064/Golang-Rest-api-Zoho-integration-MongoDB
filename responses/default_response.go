
package responses

type DefaultResponse struct {
    Status  int                    `json:"status"`
    Message string                 `json:"message"`
    Data    interface{}            `json:"data"`
}

type ErrorResponse struct {
    Status  int                    `json:"status"`
    Message string                 `json:"message"`
    Data    interface{}            `json:"error"`
}


