function calc {
	res=`echo $1 | bc -l`
	### command line test
	echo "op of $1 => $res"
	zenity --info --text "$1 = \n \n $res"
}


### some test cases

calc "2+2"
echo $?
calc "100/3"
echo $?
