/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package plugins

import (
	"configcenter/src/common"
	"configcenter/src/thirdparty/monitor/meta"
	"configcenter/src/thirdparty/monitor/plugins/blueking"
	"configcenter/src/thirdparty/monitor/plugins/noop"
)

// Plugin is interface for self-defined monitor plugin
type Plugin interface {
	Report(c meta.Content) error
}

// GetPlugin get the plugin according by the plugin name
// if no plugin is found by the name , use Noop
func GetPlugin(pluginName string) (Plugin, error) {
	switch pluginName {
	case common.BKNoopMonitorPlugin:
		return noop.NewNoop(), nil
	case common.BKBluekingMonitorPlugin:
		return blueking.NewBKmonitor()
	default:
		return noop.NewNoop(), nil
	}
}
