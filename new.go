package asynq

// srv.heartbeater.serverID
func (srv *Server) ServerID() string {
	return srv.heartbeater.serverID
}

// srv.state.value
func (srv *Server) State() serverStateValue {
	// serverStateValue 状态码改大写
	return srv.state.value
}

// 返回队列ID
func (w *ResultWriter) QueueID() string { return w.qname }

// 增加到所有redis.Option.ClientName中
// 正则查找 redis..*option
const ClientName string = "asynq"
