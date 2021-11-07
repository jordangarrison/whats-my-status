PACKAGE_NAME=wms
PACKAGE_VERSION=0.1.0
PACKAGE_CONFIG_TYPE=yaml

build:
	bash ./buildscript.sh ${PACKAGE_NAME} ${PACKAGE_VERSION}
fast-build:
	bash ./buildscript.sh ${PACKAGE_NAME} ${PACKAGE_VERSION} fast
clean:
	rm -r ./build