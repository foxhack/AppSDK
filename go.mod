module github.com/edgexfoundry/nsplussdk

require (
	bitbucket.org/bertimus9/systemstat v0.0.0-20180207000608-0eeff89b0690
	github.com/BurntSushi/toml v0.3.1
	github.com/eclipse/paho.mqtt.golang v1.1.1
	github.com/edgexfoundry/go-mod-core-contracts v0.0.0-20190508195957-b02b9d963a25
	github.com/edgexfoundry/go-mod-messaging v0.0.0-20190327144236-4222ae1edb0b
	github.com/edgexfoundry/go-mod-registry v0.0.0-20190401195203-552208258719
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/go-kit/kit v0.8.0
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/google/uuid v1.1.0
	github.com/gorilla/context v1.1.1
	github.com/gorilla/mux v1.7.0
	github.com/imdario/mergo v0.3.6
	github.com/magiconair/properties v1.8.0
	github.com/mattn/go-xmpp v0.0.0-20190124093244-6093f50721ed
	github.com/pebbe/zmq4 v1.0.0
	github.com/pelletier/go-toml v1.2.0
	github.com/pkg/errors v0.8.1
	github.com/robfig/cron v0.0.0-20180505203441-b41be1df6967
	github.com/stretchr/testify v1.3.0
	github.com/tealeg/xlsx v1.0.5
	github.com/thinkeridea/go-extend v1.1.0
	github.com/ugorji/go/codec v0.0.0-20190316192920-e2bddce071ad
	golang.org/x/text v0.3.0
	gopkg.in/eapache/queue.v1 v1.1.0
	gopkg.in/mgo.v2 v2.0.0
	gopkg.in/yaml.v2 v2.2.2
)

replace (
	cloud.google.com/go v0.26.0 => github.com/googleapis/google-cloud-go v0.26.0
	github.com/edgexfoundry/go-mod-core-contracts v0.0.0-20190508195957-b02b9d963a25 => /home/hgf/golang/src/github.com/edgexfoundry/go-mod-core-contracts
	github.com/edgexfoundry/go-mod-messaging v0.0.0-20190327144236-4222ae1edb0b => /home/hgf/golang/src/github.com/edgexfoundry/go-mod-messaging
	golang.org/x/crypto v0.0.0-20180505025534-4ec37c66abab => github.com/golang/crypto v0.0.0-20180505025534-4ec37c66abab
	golang.org/x/crypto v0.0.0-20180904163835-0709b304e793 => github.com/golang/crypto v0.0.0-20180904163835-0709b304e793
	golang.org/x/crypto v0.0.0-20181029021203-45a5f77698d3 => github.com/golang/crypto v0.0.0-20181029021203-45a5f77698d3
	golang.org/x/crypto v0.0.0-20181203042331-505ab145d0a9 => github.com/golang/crypto v0.0.0-20181203042331-505ab145d0a9
	golang.org/x/exp v0.0.0-20180321215751-8460e604b9de => github.com/golang/exp v0.0.0-20180321215751-8460e604b9de
	golang.org/x/exp v0.0.0-20181112044915-a3060d491354 => github.com/golang/exp v0.0.0-20181112044915-a3060d491354
	golang.org/x/lint => github.com/golang/lint v0.0.0-20181026193005-c67002cb31c3
	golang.org/x/net v0.0.0-20180724234803-3673e40ba225 => github.com/golang/net v0.0.0-20180724234803-3673e40ba225
	golang.org/x/net v0.0.0-20180826012351-8a410e7b638d => github.com/golang/net v0.0.0-20180826012351-8a410e7b638d
	golang.org/x/net v0.0.0-20180906233101-161cd47e91fd => github.com/golang/net v0.0.0-20180906233101-161cd47e91fd
	golang.org/x/net v0.0.0-20181023162649-9b4f9f5ad519 => github.com/golang/net v0.0.0-20181023162649-9b4f9f5ad519
	golang.org/x/net v0.0.0-20181201002055-351d144fa1fc => github.com/golang/net v0.0.0-20181201002055-351d144fa1fc
	golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be => github.com/golang/oauth2 v0.0.0-20180821212333-d2e6202438be
	golang.org/x/oauth2 v0.0.0-20181017192945-9dcd33a902f4 => github.com/golang/oauth2 v0.0.0-20181017192945-9dcd33a902f4
	golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f => github.com/golang/sync v0.0.0-20180314180146-1d60e4601c6f
	golang.org/x/sync v0.0.0-20181221193216-37e7f081c4d4 => github.com/golang/sync v0.0.0-20181221193216-37e7f081c4d4
	golang.org/x/sys v0.0.0-20180823144017-11551d06cbcc => github.com/golang/sys v0.0.0-20180823144017-11551d06cbcc
	golang.org/x/sys v0.0.0-20180830151530-49385e6e1522 => github.com/golang/sys v0.0.0-20180830151530-49385e6e1522
	golang.org/x/sys v0.0.0-20180903190138-2b024373dcd9 => github.com/golang/sys v0.0.0-20180903190138-2b024373dcd9
	golang.org/x/sys v0.0.0-20180905080454-ebe1bf3edb33 => github.com/golang/sys v0.0.0-20180905080454-ebe1bf3edb33
	golang.org/x/sys v0.0.0-20180906133057-8cf3aee42992 => github.com/golang/sys v0.0.0-20180906133057-8cf3aee42992
	golang.org/x/sys v0.0.0-20180909124046-d0be0721c37e => github.com/golang/sys v0.0.0-20180909124046-d0be0721c37e
	golang.org/x/sys v0.0.0-20181026203630-95b1ffbd15a5 => github.com/golang/sys v0.0.0-20181026203630-95b1ffbd15a5
	golang.org/x/sys v0.0.0-20181030150119-7e31e0c00fa0 => github.com/golang/sys v0.0.0-20181030150119-7e31e0c00fa0
	golang.org/x/sys v0.0.0-20181205085412-a5c9d58dba9a => github.com/golang/sys v0.0.0-20181205085412-a5c9d58dba9a
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
	golang.org/x/time v0.0.0-20180412165947-fbb02b2291d2 => github.com/golang/time v0.0.0-20180412165947-fbb02b2291d2
	golang.org/x/tools v0.0.0-20180525024113-a5b4c53f6e8b => github.com/golang/tools v0.0.0-20180525024113-a5b4c53f6e8b
	golang.org/x/tools v0.0.0-20180828015842-6cd1fcedba52 => github.com/golang/tools v0.0.0-20180828015842-6cd1fcedba52
	golang.org/x/tools v0.0.0-20181221154417-3ad2d988d5e2 => github.com/golang/tools v0.0.0-20181221154417-3ad2d988d5e2
	google.golang.org/api v0.0.0-20181021000519-a2651947f503 => github.com/googleapis/google-api-go-client v0.0.0-20181021000519-a2651947f503
	google.golang.org/appengine v1.1.0 => github.com/golang/appengine v1.1.0
	google.golang.org/appengine v1.2.0 => github.com/golang/appengine v1.2.0
	google.golang.org/appengine v1.4.0 => github.com/golang/appengine v1.4.0
	google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8 => github.com/google/go-genproto v0.0.0-20180817151627-c66870c02cf8
	google.golang.org/genproto v0.0.0-20181016170114-94acd270e44e => github.com/google/go-genproto v0.0.0-20181016170114-94acd270e44e
	google.golang.org/grpc v1.14.0 => github.com/grpc/grpc-go v1.14.0
	google.golang.org/grpc v1.15.0 => github.com/grpc/grpc-go v1.15.0
	gopkg.in/mgo.v2 v2.0.0 => gopkg.in/mgo.v2 v2.0.0-20180705113604-9856a29383ce
)
