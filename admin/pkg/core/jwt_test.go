package core

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTokenCreate(t *testing.T) {
	option := CustomClaimsOption{
		AccountId:     1,
		AccountName:   "admin",
		ExpireSeconds: 3600,
		UUID:          1,
		Secret:        "MIIEowIBAAKCAQEA6k9lyCPE/s0A4XU0InSJqKXmN3cv3lhcX+nIFGo1YyHfDUOoxhX3qATaFft0pUFgkpoUVXIGYXRL+3Th+S/ilI6kCDaq4dQrL2EppT8JsvdfpOWJYa4eLbl0x5mUYSWc0Z50vZgeE66oV6ZCfd31yV+kMK1vhE4NpuhkaeDZwl/RZWyn39qlXaUnDuiKix/gFqxDmOzgLNa1bNcHDesfhMo+uOF4sXz/UsaqecSHpjC6WE1l7CIE5VHRjdCnl/uTg0JIJxl3U2JAfa50zjhFtfvrsLvte4yVTZPzBPsxKam6MX/C6B7aE6YyY7j2Pyq+EsiV0ofqmpRbn1M34veWrwIDAQABAoIBAGr/5xMGgqd5JrXOuAgj/9ksKX2ayBlZcJW8RMpN1iN3DE0aYLBUx8vKq2zub84nlNpd0ntSnudSHICwV6Fb5LTCjtXRYSfPpAj27fWwW9WGcf7zc1FQGif3UQT1dXUHoB15pwYNuW5zjBy6qBKikeNr8abpDMv2ePWHmpNfQxInn81ONDLiWXN+c2j2bwoP6TmkRag6WrZMjyCyKeXbZSb6eiu0J2d20XeV6k86LsqH5BNXjvW91nu1b8W8GqUgqMx4qYeUJy2RkFiniZLUWjt18nNy1W7SKkbUjBJRGWoSp22fz6WR7WXDdKfK7oo+SPL1A7YBoZBLdzXVbdiHJ4ECgYEA+zwptzW0+5Oj5t22ajFOmgkR5EeGVGNlGSDTa0I6jybTwuj3/skV5wwh5E8hDP1uN0vhiYyUMPbY+ChnqgP+jFsn52t5GWcrKwjrjsgJSHss1H+TQU8fmntF42sX8fPZ1Scpw8cmaxTdgqCiUX+PFWU+8q/sYj6XelKyr0sl9QkCgYEA7sEO63Rfc0SiRy1ZEEPttnjJhvgqpNO28/r7GdhfYQ+c1E8SKn89WBE0U26wXHDc8CsMGzdGjh5ome790xJroUmfJqINl950R2RFPYbreRdcNoV9LVtfa8rbJADU+zsu7q/O7/Wj5WIjZ762YCLtriRIXZVjj1KjYG7Nh5WEk/cCgYEAmCegsBezy2VoL2r5jijcjMePcWyr+zHTSbuqr1wF8Sq3t2S2xTTtseeUIxyVTLOz64NmTZK5MUperzRxS/NtV33hlfNt/SeRalfVqbyWerI2vV+iRDxyTHH6KrhyYROpSsGUNsDHELUrTnCQvcD+XBvcWsiW5g04Wln5Y9akX1kCgYBDCDEIKkQiOMLW+QRobKnw8TijjElsvYRAO9cECpu8xIVjlk+jXtX6Q9pNEmaxjX56b5uWiI5mmGMF1POurx1Iy0LDtfvbB7nd0WkfPSIffU+GfzQF6AGCYhHkqdciMhubKAERngXGASevmPvbQfM5UTaAD9FEDA+So4lQ58mxQQKBgFHwz/dgVuAlCH/bx0ife2GoBFhzost94LOzr7jvGmdTCqvmiE/Pf3qbkviCJjhGMAvZkTiLiZmKY3ddCz6GVQAxzrOTylOBxDmydbUFr5vjANEvRHoMusJoJOWF8LYqzUlByq83QnbJvxRrIYhnSV8SJ+VAkRITVD3coWUOXHXf",
	}
	token, err := JWTCreate(option)
	assert.Nil(t, err, "错误")
	fmt.Println("token: ", token)

	res, err := JWTCheck(token, option.Secret)
	assert.Nil(t, err, "错误")
	fmt.Println(res.ID, res.AdminName, res.AdminID, res.Subject, res)
}
