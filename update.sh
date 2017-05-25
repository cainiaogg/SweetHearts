# dirname $0
cc=`pwd`
# echo $0 $cc $1

Gruation=$1
echo start update $Gruation

Gruation_old=""$Gruation".old"

rm -rf $Gruation_old
echo "rm -rf $Gruation_old"

mv $Gruation $Gruation_old
echo "mv $Gruation $Gruation_old"

cp -r Gruation $Gruation
echo "cp -r Gruation $Gruation"

run="$cc/$Gruation/run.sh"
sed -i "s/SweetHearts/$Gruation/g" $run

main="$cc/$Gruation/main.go"
sed -i "s/beego.BConfig.Listen.AdminPort = 8088/beego.BConfig.Listen.AdminPort = $3/g" $main
sed -i "s/beego.Run()/beego.Run($2)/g" $main
