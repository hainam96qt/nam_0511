package attendance

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"nam_0511/internal/repo/contracts"
	"nam_0511/pkg/smartcontract"
	"nam_0511/pkg/util/random"
)

type Repository struct {
	ethClient *ethclient.Client
	config    smartcontract.SmartContractConfig
}

func NewPostgresRepository(ethClient *ethclient.Client, config smartcontract.SmartContractConfig) *Repository {
	return &Repository{
		ethClient: ethClient,
		config:    config,
	}
}

func (r *Repository) CheckIn(ctx context.Context, req *contracts.AttendanceContractAttendanceRecord) error {
	contractAddress := common.HexToAddress(r.config.ContractAddress)

	_, err := r.ethClient.CodeAt(context.Background(), contractAddress, nil)
	if err != nil {
		return err
	}

	contract, err := contracts.NewContracts(contractAddress, r.ethClient)
	if err != nil {
		return err
	}

	privateKey, err := crypto.HexToECDSA(r.config.PrivateKey)
	if err != nil {
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
	if err != nil {
		return err
	}

	tx, err := contract.RecordAttendance(auth, req.EmployeeId, req.CheckInTime, req.Details)
	if err != nil {
		return err
	}

	// quan ly giao dich
	_ = tx

	return nil
}

func (r *Repository) GetListAttendanceByUserID(ctx context.Context, userID int64) ([]contracts.AttendanceContractAttendanceRecord, error) {
	contractAddress := common.HexToAddress(r.config.ContractAddress)

	_, err := r.ethClient.CodeAt(context.Background(), contractAddress, nil)
	if err != nil {
		return nil, err
	}

	contract, err := contracts.NewContracts(contractAddress, r.ethClient)
	if err != nil {
		return nil, err
	}
	data, err := contract.GetAttendanceByEmployeeId(nil, big.NewInt(userID))
	if err != nil {
		log.Fatal(err)
	}

	return data, nil
}

func (r *Repository) AuthorizeEntity(ctx context.Context) (string, error) {
	contractAddress := common.HexToAddress(r.config.ContractAddress)

	contract, err := contracts.NewContracts(contractAddress, r.ethClient)
	if err != nil {
		log.Fatal(err)
	}

	userPrivateHexKey, err := random.GenerateHex64()
	if err != nil {
		return "", err
	}

	userPrivateKey, err := crypto.HexToECDSA(userPrivateHexKey)
	if err != nil {
		return "", err
	}

	privateKey, err := crypto.HexToECDSA(r.config.PrivateKey)
	if err != nil {
		log.Fatal("private key not found", err)
	}

	sender := crypto.PubkeyToAddress(userPrivateKey.PublicKey)

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
	if err != nil {
		log.Fatal(err)
	}

	tx, err := contract.AuthorizeEntity(auth, sender)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx.Data())

	return userPrivateHexKey, err
}
