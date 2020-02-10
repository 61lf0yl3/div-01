#!/bin/bash
echo "Expected output:"
cat 1.txt
echo ""
echo "Actual output:"
./ascii-art-fs "hello" standard
echo "*************************************************************************************************************************"

echo "Expected output:"
cat 2.txt
echo ""
echo "Actual output:"
./ascii-art-fs "hello world" shadow
echo "*************************************************************************************************************************"

echo "Expected output:"
cat 3.txt
echo ""
echo "Actual output:"
./ascii-art-fs "nice 2 meet you" thinkertoy
echo "*************************************************************************************************************************"

echo "Expected output:"
cat 4.txt
echo ""
echo "Actual output:"
./ascii-art-fs "you & me" standard
echo "*************************************************************************************************************************"

echo "Expected output:"
cat 5.txt
echo ""
echo "Actual output:"
./ascii-art-fs "123" shadow
echo "*************************************************************************************************************************"


echo "Expected output:"
cat 6.txt
echo ""
echo "Actual output:"
./ascii-art-fs "/(\")" thinkertoy
echo "*************************************************************************************************************************"


echo "Expected output:"
cat 7.txt
echo ""
echo "Actual output:"
./ascii-art-fs ABCDEFGHIJKLMNOPQRSTUVWXYZ shadow
echo "*************************************************************************************************************************"


echo "Expected output:"
cat 8.txt
echo ""
echo "Actual output:"
./ascii-art-fs "\"#$%&/()*+,-./" thinkertoy
echo "*************************************************************************************************************************"


echo "Expected output:"
cat 9.txt
echo ""
echo "Actual output:"
./ascii-art-fs "It's Working" thinkertoy
echo "*************************************************************************************************************************"


