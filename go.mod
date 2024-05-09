module github.com/mises-id/mises-vpnsvc

go 1.18

require (
	github.com/caarlos0/env v3.5.0+incompatible
	github.com/ethereum/go-ethereum v1.12.0
	github.com/go-kit/kit v0.13.0
	github.com/gogo/protobuf v1.3.2
	github.com/gorilla/mux v1.8.0
	github.com/jinzhu/inflection v1.0.0
	github.com/joho/godotenv v1.5.1
	github.com/metaverse/truss v0.3.1
	github.com/mises-id/mises-miningsvc v0.0.0-20230908061632-02cbf7c7bfd8
	github.com/mises-id/sns-storagesvc/sdk v0.0.0-20220920081129-d682f954bf94
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.9.3
	github.com/syndtr/goleveldb v1.0.1-0.20220721030215-126854af5e6d
	go.mongodb.org/mongo-driver v1.12.1
	google.golang.org/grpc v1.57.0
)

require (
	github.com/StackExchange/wmi v0.0.0-20180116203802-5d049714c4a6 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.3.2 // indirect
	github.com/deckarep/golang-set/v2 v2.1.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.1.0 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golang/snappy v0.0.5-0.20220116011046-fa5810519dcb // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/holiman/uint256 v1.2.3 // indirect
	github.com/klauspost/compress v1.16.3 // indirect
	github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe // indirect
	github.com/shirou/gopsutil v3.21.4-0.20210419000835-c7a38de76ee5+incompatible // indirect
	github.com/tklauser/go-sysconf v0.3.10 // indirect
	github.com/tklauser/numcpus v0.4.0 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	github.com/zksync-sdk/zksync2-go v0.3.1
	golang.org/x/crypto v0.12.0 // indirect
	golang.org/x/exp v0.0.0-20230810033253-352e893a4cad // indirect
	golang.org/x/net v0.14.0 // indirect
	golang.org/x/sync v0.3.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
	golang.org/x/text v0.12.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230807174057-1744710a1577 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce // indirect
)

replace github.com/go-kit/kit => github.com/mises-id/kit v0.12.1-0.20211203081751-bc5397e8a165

replace go.mongodb.org/mongo-driver => go.mongodb.org/mongo-driver v1.11.6
