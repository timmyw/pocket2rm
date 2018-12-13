#!/bin/sh

pdfgen=$1 # wkhtmltopdf full path
src_url=$2
output=$3

libdir=$(dirname $pdfgen)

pdfopts=--grayscale

LD_LIBRARY_PATH=$LD_LIBRARY_PATH:$libdir $pdfgen $pdfopts $src_url $output
