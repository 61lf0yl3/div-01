#!/bin/bash
echo "Expected output:"
cat 1.txt
echo ""
echo "Actual output:"
./ascii-art-output "First\nTest" shadow --output=test00.txt
cat test00.txt
echo "***********************************************************************************"

echo "Expected output:"
cat 2.txt
echo ""
echo "Actual output:"
./ascii-art-output "hello" standard --output=test01.txt
cat test01.txt
echo "***********************************************************************************"

echo "Expected output:"
cat 3.txt
echo ""
echo "Actual output:"
./ascii-art-output "123 -> #$%" standard --output=test02.txt
cat test02.txt
echo "***********************************************************************************"

echo "Expected output:"
cat 4.txt
echo ""
echo "Actual output:"
./ascii-art-output "432 -> #$%&@" shadow --output=test03.txt
cat test03.txt
echo "***********************************************************************************"

echo "Expected output:"
cat 5.txt
echo ""
echo "Actual output:"
./ascii-art-output "There" shadow --output=test04.txt
cat test04.txt
echo "***********************************************************************************"


echo "Expected output:"
cat 6.txt
echo ""
echo "Actual output:"
./ascii-art-output "123 -> \"#$%@" thinkertoy --output=test05.txt
cat test05.txt
echo "***********************************************************************************"


echo "Expected output:"
cat 7.txt
echo ""
echo "Actual output:"
./ascii-art-output "2 you" thinkertoy --output=test06.txt
cat test06.txt
echo "***********************************************************************************"


echo "Expected output:"
cat 8.txt
echo ""
echo "Actual output:"
./ascii-art-output "Testing long output!" standard --output=test07.txt
cat test07.txt
echo "***********************************************************************************"
