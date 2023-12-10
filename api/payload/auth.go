package payload

type GenJwtReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GenJwtResp struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

type ValidJwtReq struct {
	Token string `json:"token"`
}

type ValidJwtResp struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}
