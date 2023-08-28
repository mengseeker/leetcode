package rbac

import "leetcode/rbac/api"

type server struct {
	api.UnimplementedAgentServer
}

func (s server) HasPermission(req *api.PermissionReq, stream api.Agent_HasPermissionServer) error {
	
	return nil
}
