package api

func (m *Methods) SuccessResponse(msg string, data interface{}) *Methods {
	m.resp.Success = true
	m.resp.Message = msg
	m.resp.Data = data

	return m
}

func (m *Methods) BadRequest(err error) *Methods {
	m.resp.Success = false
	m.resp.Message = err.Error()

	return m
}

func (m *Methods) MethodNotFound() *Methods {
	m.resp.Success = false
	m.resp.Message = "Method not found"

	return m
}
