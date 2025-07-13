//go:build !solution

package retryupdate

import (
	"errors"

	"github.com/gofrs/uuid"
	"gitlab.com/slon/shad-go/retryupdate/kvapi"
)

func getValueWithRetries(c kvapi.Client, key string) (value *string, version *uuid.UUID, err error) {
	getReq := &kvapi.GetRequest{Key: key}

	for {
		getResp, getErr := c.Get(getReq)

		if getErr == nil {
			return &getResp.Value, &getResp.Version, getErr
		}

		var authErr *kvapi.AuthError
		if errors.As(getErr, &authErr) {
			return nil, nil, getErr
		}

		if errors.Is(getErr, kvapi.ErrKeyNotFound) {
			return nil, nil, nil
		}
	}
}

func UpdateValue(c kvapi.Client, key string, updateFn func(oldValue *string) (newValue string, err error)) error {
getLoop:
	for {
		getValue, oldVersion, getErr := getValueWithRetries(c, key)
		if getErr != nil {
			return getErr
		}

	updateLoop:
		for {
			newValue, updateErr := updateFn(getValue)

			if updateErr != nil {
				return updateErr
			}

			setReq := &kvapi.SetRequest{Key: key, Value: newValue}
			if oldVersion != nil {
				setReq.OldVersion = *oldVersion
			}

			for {
				_, setErr := c.Set(setReq)

				var authErr *kvapi.AuthError
				if setErr == nil || errors.As(setErr, &authErr) {
					return setErr
				}

				if errors.Is(setErr, kvapi.ErrKeyNotFound) {
					getValue = nil
					oldVersion = nil
					continue updateLoop
				}

				var conflictError *kvapi.ConflictError
				if errors.As(setErr, &conflictError) {
					if conflictError.ExpectedVersion == setReq.NewVersion {
						return nil
					}
					continue getLoop
				}
			}
		}
	}
}
