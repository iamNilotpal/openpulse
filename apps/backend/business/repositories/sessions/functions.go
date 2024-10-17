package sessions

import sessions_store "github.com/iamNilotpal/openpulse/business/repositories/sessions/store/postgres"

func ToNewDBSession(cmd NewSession) sessions_store.NewSession {
	return sessions_store.NewSession{
		Token:        cmd.Token,
		UserId:       cmd.UserId,
		UserAgent:    cmd.UserAgent,
		IpAddress:    cmd.IpAddress,
		DeviceInfo:   cmd.DeviceInfo,
		LocationInfo: cmd.LocationInfo,
	}
}
