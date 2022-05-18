#Consultar um ativo específico no razão
source variaveis_de_ambienteOrg1.sh
peer chaincode query -C mychannel -n contratointeligente -c '{"Args":["ConsultarNotificacao","2"]}'
