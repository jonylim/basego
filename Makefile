-include .env

DATE := $(shell date '+%Y%m%d')
APP_NAME := basego-api
EXEC_NAME := basego-api
APIDOC_NAME := apidoc-basego-api
DEST_DIR := /opt/basego/basego-api
BACKUP_DIR := ${DEST_DIR}/backup/${DATE}

all: clean build

dep:
	@echo "Installing missing dependencies..."
	dep ensure

apidoc:
	apidoc -i internal/app/${APP_NAME} -o ${APIDOC_NAME};

build:
	go build -o ${EXEC_NAME} cmd/${APP_NAME}/main.go

build-exe:
	go build -o ${EXEC_NAME}.exe cmd/${APP_NAME}/main.go

build-for-linux:
	GOOS=linux GOARCH=amd64 go build -o ${EXEC_NAME} cmd/${APP_NAME}/main.go

clean:
	@if [ -f ${EXEC_NAME} ]; then\
		rm ${EXEC_NAME};\
	fi

deploy:
	@if [ ! -f ${EXEC_NAME} ]; then\
		echo "Run \`make <app> build\` first!";\
		exit 1;\
	fi

	@if [ ! -d ${DEST_DIR} ]; then\
		sudo mkdir -pv ${DEST_DIR};\
	fi

	@if [ -f ${DEST_DIR}/init/${EXEC_NAME}.service ]; then\
		echo "sudo systemctl stop ${EXEC_NAME}.service";\
		sudo systemctl stop ${EXEC_NAME}.service;\
	fi

	@echo $(shell date) ${APP_NAME} >> deploy_history

	@echo Deploying to ${DEST_DIR}

	@sudo cp ${EXEC_NAME} ${DEST_DIR}/
	@if [ -d ${APIDOC_NAME} ]; then\
		sudo cp -r ${APIDOC_NAME} ${DEST_DIR}/;\
	fi
	@sudo cp -r assets ${DEST_DIR}/
	@sudo cp -r init ${DEST_DIR}/
	@sudo cp -r web ${DEST_DIR}/
	
	@echo Copying backup to ${BACKUP_DIR}

	@sudo mkdir -pv ${BACKUP_DIR}
	@sudo cp ${EXEC_NAME} ${BACKUP_DIR}/
	@if [ -d ${APIDOC_NAME} ]; then\
		sudo cp -r ${APIDOC_NAME} ${BACKUP_DIR}/;\
	fi
	@sudo cp -r assets ${BACKUP_DIR}/
	@sudo cp -r init ${BACKUP_DIR}/
	@sudo cp -r web ${BACKUP_DIR}/

#	@echo "sudo systemctl start ${EXEC_NAME}.service"
	sudo systemctl start ${EXEC_NAME}.service

restart-service:
	sudo systemctl stop ${EXEC_NAME}.service
	sudo systemctl start ${EXEC_NAME}.service

.PHONY: all dep apidoc install build build-exe build-for-linux clean deploy help