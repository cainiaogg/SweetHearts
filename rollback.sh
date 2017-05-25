Graution=$1
Graution_old="$Graution.old"
rm -rf $Graution
cp -r $Graution_old $Graution
echo "cp -r Graution_old Graution"
killall $Graution
