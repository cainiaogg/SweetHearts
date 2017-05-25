# dirname $0
cc=`pwd`
# echo $0 $cc $1

Graution=$1
BeeGo=$2
AdminBeeGo=$3
echo start update $Graution

Graution_old=""$Graution".old"

rm -rf $Graution_old
echo "rm -rf $Graution_old"

mv $Graution $Graution_old
echo "mv $Graution $Graution_old"

cp -r Graution $Graution
echo "cp -r Graution $Graution"

run="$cc/$Graution/src/SweetHearts/run.sh"
sed -i "s/SweetHearts/$Graution/g" $run

main="$cc/$Graution/src/SweetHearts/main.go"
sed -i "s/beego.BConfig.Listen.AdminPort = 8088/beego.BConfig.Listen.AdminPort = $AdminBeeGo/g" $main
sed -i "s/beego.Run()/beego.Run("$BeeGo")/g" $main

killall $Graution
