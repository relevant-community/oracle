# Makefile for the "oraclenode" docker image.

all:
	docker build --tag relevant/oracle oraclenode

.PHONY: all