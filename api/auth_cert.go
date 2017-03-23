package api

// CertAuth is used to perform cert backend operations on Vault
type CertAuth struct {
	c *Client
}

// Token is used to return the client for token-backend API calls
func (a *Auth) Cert() *CertAuth {
	return &CertAuth{c: a.c}
}

func (c *CertAuth) Login() (*Secret, error) {
	r := c.c.NewRequest("POST", "/v1/auth/cert/login")

	resp, err := c.c.RawRequest(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ParseSecret(resp.Body)
}
