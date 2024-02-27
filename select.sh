#!/usr/bin/env bash

set -eo pipefail

_year=${YEAR:-$(($(date +"%Y")+1))}
_passes=2
_cfg=${CONFIG:-"cfg/base.yaml,cfg/rm2.base.yaml,cfg/template_months_on_side.yaml,cfg/rm2.mos.default.yaml,cfg/rm2.mos.default.dailycal.yaml"}
_name=${NAME:-"rm2.mos.default.dailycal"}
_preview=${PREVIEW}
sed -i 's/dotted: true/dotted: false/' cfg/base.yaml

echo "Starting ${_year} Planner"
sleep 1s

PLANNER_YEAR="${_year}" PASSES="${_passes}" CFG="${_cfg}" NAME="${_name}.${_year}" PREVIEW="${_preview}" ./single.sh