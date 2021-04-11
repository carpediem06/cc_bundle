#     			-*- Autoconf -*-
# Process this file with autoconf to produce a configure script.

# Copyright (C) 1992-1995, 1999-2017, 2020-2021 Free Software
# Foundation, Inc.

# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program.  If not, see <https://www.gnu.org/licenses/>.

AC_INIT([app], [1.0], [ma.06.am@hotmail.com], [carpediem06], [http://www.github.com/carpediem06])

AC_PROG_CC
AM_CONFIG_HEADER(config.h)
AC_CONFIG_MACRO_DIR([m4])
m4_ifdef([AM_SILENT_RULES], [AM_SILENT_RULES([yes])])


dnl -----------------------------------------------
dnl PACKAGE_NAME & PACKAGE_VERSION
dnl -----------------------------------------------

GENERIC_PACKAGE_NAME=app

PACKAGE=$GENERIC_PACKAGE_NAME
AC_SUBST(GENERIC_PACKAGE_NAME)

GENERIC_MAJOR_VERSION=1
GENERIC_MINOR_VERSION=0

GENERIC_VERSION=$GENERIC_MAJOR_VERSION.$GENERIC_MINOR_VERSION
GENERIC_RELEASE=$GENERIC_MAJOR_VERSION.$GENERIC_MINOR_VERSION
AC_SUBST(GENERIC_RELEASE)
AC_SUBST(GENERIC_VERSION)
VERSION=$GENERIC_VERSION

AM_INIT_AUTOMAKE($PACKAGE, $VERSION, no-define)

AC_PROG_CC
AM_SANITY_CHECK
AX_PROG_SPLINT(["+posixlib +unixlib +gnuextensions -formatconst -nullpass -UG_HAVE_ISO_VARARGS -D__gnuc_va_list=va_list"])

dnl -----------------------------------------------
dnl Generates Makefile
dnl -----------------------------------------------

AC_OUTPUT(Makefile \
        app/Makefile \
)

