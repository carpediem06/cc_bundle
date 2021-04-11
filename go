#!/bin/bash
##################################################
# File     : go
# Purpose  : compilation rules
##################################################

export R_COLOR="\033[1;31;40m"
export Y_COLOR="\033[1;33m"
export G_COLOR="\033[1;32m"
export NO_COLOR="\033[0m"
export AUTODIR=auto

#DoCmake: build cmake modules
function DoCmake {
  for item in ./*
   do
     if [[ -d "$item" ]]; then
        if [[ -f $item/CMakeLists.txt ]]; then
	  echo -e "${Y_COLOR}cmake: $item${NO_COLOR}"
	  cd $item	
	  mkdir -p build; CheckCmd
	  cmake .; CheckCmd
	  make ${MAKEOPT}; CheckCmd
	  cd ..
        else
	  cd $item
	     DoCmake
          cd ..
	fi
     fi
   done
}

# DoRecursively: "copy" or "clean" or "build"
function DoRecursively {
    action=$1
    where=$2
    find=0	
    #echo -e "action: $action"
    #echo -e "where: $where"
    for item in ./*
    do
      if [[ -d "$item" ]]; then
	# Not empty dir
    	if [[ ! -z "$(ls -A $item)" ]] && [[ "$item" != "./${AUTODIR}" ]]; then
	  if [[ $action == "clean" ]]; then
	    # Check Makefile.in
    	    if [[ -f "$item/Makefile.am" ]]; then
    		echo -e "Clean: $item"
                rm -f $item/Makefile.in
    	        if [[ -d "$item/CMakeFiles" ]]; then
		 echo -e "Clean: $item/CMakeFiles"
		 rm -rf $item/CMakeFiles
		fi
		if [[ -d "$item/build" ]]; then
		  echo -e "Clean: $item/build"
		  rm -rf $item/build
		fi
    		test -f Makefile && make clean && make distclean
    		rm -f $item/Makefile
    		rm -rf $item/.deps
    		
	    else  
		cd $item
		DoRecursively $action $where
		cd ..
    	    fi
          elif [[ $action == "build" ]]; then
	    if [[ -f "$item/CMakeCache.txt" ]]; then
		echo -e "CMAKE: $item"
		mkdir -p $item/build
	    else
		cd $item	
		DoRecursively $action $where
		cd ..
	    fi
	  else # copy 
	    if [[ -f "$item/Makefile.am" ]]; then
		echo -e " $item"
		cp $where/${AUTODIR}/Makefile.in $item
		find=1
	    else
		cd $item	
		DoRecursively $action $where
		cd ..
	    fi
	  fi
    	fi
      fi
    done
    return $find
}

# CheckCmd: detect error
function CheckCmd {
    if [ "$?" -ne "0" ]; then
	echo -e "Exit with ${R_COLOR}error${NO_COLOR}\n"
        exit 1;
    fi
}

if [[ $1 == "clean" ]]; then
 echo -e "${G_COLOR}Cleaning...${NO_COLOR}"

 DoRecursively "clean" `pwd`
 find . -iwholename '*cmake*' -not -name CMakeLists.txt -delete
 rm -rf autom4te.cache \
   	m4
 rm -f AUTHORS \
      COPYING \
      INSTALL \
      NEWS \
      README \
      mkinstalldirs \
      configure \
      autogen.sh \
      Makefile \
      Makefile.in \
      config.h.in* \
      depcomp \
      missing \
      ChangeLog \
      ltmain.sh \
      install-sh \
      aclocal.m4 \
      compile \
      config.* \
      aminclude_static.am
      
else

 echo -e "${G_COLOR}Compiling...${NO_COLOR}"

 MKOPT="-j 8"
 COMPILEAT=$(date +"%H:%M")
 ROOTFS=../rootfs
 LINUX=../linux
 TFTPDIR=/tftpboot
 CMAKE=0

if [[ $CMAKE -eq 1 ]];then

  DoCmake
  
else
#  . cc_buildroot_toolchain
  . cc_arm_toolchain
#  . cc_x86_64_toolchain

  DoRecursively "copy" `pwd`
  cp -r ${AUTODIR}/* .
  autoreconf -f -i; CheckCmd
  ./autogen.sh --host=${HOST}; CheckCmd
  ./configure --prefix=${PREFIX}/usr --host=${HOST} --build=${BUILD}; CheckCmd
  make ${MKOPT}; CheckCmd
  
fi


echo ""
echo ""
echo "Done"

fi

