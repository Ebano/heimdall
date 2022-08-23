package keystore

import (
	"crypto/rand"
	"crypto/rsa"

	"github.com/rs/zerolog"
	"go.uber.org/fx"

	"github.com/dadrus/heimdall/internal/config"
	"github.com/dadrus/heimdall/internal/heimdall"
	"github.com/dadrus/heimdall/internal/x/errorchain"
)

// nolint
var Module = fx.Options(
	fx.Provide(NewKeyStore),
)

func NewKeyStore(conf config.Configuration, logger zerolog.Logger) (KeyStore, error) {
	const rsa2048 = 2048

	var (
		ks  KeyStore
		err error
	)

	if len(conf.Signer.KeyStore) == 0 {
		logger.Warn().Msg("Key store is not configured. NEVER DO IT IN PRODUCTION!!!! Generating an RSA key pair.")

		var privateKey *rsa.PrivateKey

		privateKey, err = rsa.GenerateKey(rand.Reader, rsa2048)
		if err != nil {
			return nil, errorchain.NewWithMessage(heimdall.ErrInternal,
				"failed to generate RSA-2048 key pair").CausedBy(err)
		}

		ks, err = NewKeyStoreFromKey(privateKey)
	} else {
		ks, err = NewKeyStoreFromPEMFile(conf.Signer.KeyStore, conf.Signer.Password)
	}

	if err != nil {
		return nil, err
	}

	logger.Info().Msg("Key store contains following entries")

	for _, entry := range ks.Entries() {
		logger.Info().Msgf("key_id: %s, algorithm: %s, size: %d", entry.KeyID, entry.Alg, entry.KeySize)
	}

	return ks, nil
}