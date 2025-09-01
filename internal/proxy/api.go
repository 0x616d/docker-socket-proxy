package proxy

const versionMatcher = `(/v[0-9.]+)?`

var api = map[string]map[string][]string{
	"ALLOW_START": {
		"POST": {`/containers/\S+/start`},
	},

	"ALLOW_STOP": {
		"POST": {`/containers/\S+/stop`},
	},

	"ALLOW_RESTART": {
		"POST": {`/containers/\S+/(stop|restart|kill)`},
	},

	"CONTAINERS": {
		"GET": {
			`/containers/json`,
			`/containers/\S+/(json|top|logs|changes|stats)`,
		},
	},

	"IMAGES": {
		"GET": {
			`/images/(json|search|get)`,
			`/images/\S+/(json|history|get)`,
		},
	},

	"BUILD": {
		"POST": {`/build(/prune)?`},
	},

	"COMMIT": {
		"POST": {`/commit`},
	},

	"NETWORKS": {
		"GET": {`/networks(/\S+)?`},
	},

	"VOLUMES": {
		"GET": {`/volumes(/\S+)?`},
	},

	"EXEC": {
		"GET": {`/exec/\S+/json`},

		"POST": {
			`/exec/\S+/(start|resize)`,
			`/containers/\S+/exec`,
		},
	},

	"SWARM": {
		"GET": {`/swarm`},
	},

	"NODES": {
		"GET": {`/nodes(/\S+)?`},
	},

	"SERVICES": {
		"GET": {
			`/services`,
			`/services/\S+(/logs)?`,
		},
	},

	"TASKS": {
		"GET": {
			`/tasks`,
			`/tasks/\S+(/logs)?`,
		},
	},

	"SECRETS": {
		"GET": {`/secrets(/\S+)?`},
	},

	"CONFIGS": {
		"GET": {`/configs(/\S+)?`},
	},

	"PLUGINS": {
		"GET": {
			`/plugins(/privileges)?`,
			`/plugins/\S+/json`,
		},
	},

	"AUTH": {
		"POST": {`/auth`},
	},

	"INFO": {
		"GET": {`/info`},
	},

	"VERSION": {
		"GET": {`/version`},
	},

	"PING": {
		"GET":  {`/_ping`},
		"HEAD": {`/_ping`},
	},

	"EVENTS": {
		"GET": {`/events`},
	},

	"SYSTEM": {
		"GET": {`/system/df`},
	},

	"DISTRIBUTION": {
		"GET": {`/distribution/\S+/json`},
	},

	"SESSION": {
		"POST": {`/session`},
	},

	"CONTAINERS_LIST":             {"GET": {`/containers/json`}},
	"CONTAINERS_CREATE":           {"POST": {`/containers/create`}},
	"CONTAINERS_INSPECT":          {"GET": {`/containers/\S+/json`}},
	"CONTAINERS_TOP":              {"GET": {`/containers/\S+/top`}},
	"CONTAINERS_LOGS":             {"GET": {`/containers/\S+/logs`}},
	"CONTAINERS_CHANGES":          {"GET": {`/containers/\S+/changes`}},
	"CONTAINERS_EXPORT":           {"GET": {`/containers/\S+/export`}},
	"CONTAINERS_STATS":            {"GET": {`/containers/\S+/stats`}},
	"CONTAINERS_RESIZE":           {"POST": {`/containers/\S+/resize`}},
	"CONTAINERS_START":            {"POST": {`/containers/\S+/start`}},
	"CONTAINERS_STOP":             {"POST": {`/containers/\S+/stop`}},
	"CONTAINERS_RESTART":          {"POST": {`/containers/\S+/restart`}},
	"CONTAINERS_KILL":             {"POST": {`/containers/\S+/kill`}},
	"CONTAINERS_UPDATE":           {"POST": {`/containers/\S+/update`}},
	"CONTAINERS_RENAME":           {"POST": {`/containers/\S+/rename`}},
	"CONTAINERS_PAUSE":            {"POST": {`/containers/\S+/pause`}},
	"CONTAINERS_UNPAUSE":          {"POST": {`/containers/\S+/unpause`}},
	"CONTAINERS_ATTACH":           {"POST": {`/containers/\S+/attach`}},
	"CONTAINERS_ATTACH_WEBSOCKET": {"GET": {`/containers/\S+/attach/ws`}},
	"CONTAINERS_WAIT":             {"POST": {`/containers/\S+/wait`}},
	"CONTAINERS_DELETE":           {"DELETE": {`/containers/\S+`}},
	"CONTAINERS_ARCHIVE_INFO":     {"HEAD": {`/containers/\S+/archive`}},
	"CONTAINERS_ARCHIVE":          {"GET": {`/containers/\S+/archive`}},
	"CONTAINERS_ARCHIVE_PUT":      {"PUT": {`/containers/\S+/archive`}},
	"CONTAINERS_PRUNE":            {"POST": {`/containers/prune`}},
	"IMAGES_LIST":                 {"GET": {`/images/json`}},
	"IMAGES_BUILD":                {"POST": {`/build`}},
	"BUILD_PRUNE":                 {"POST": {`/build/prune`}},
	"IMAGES_CREATE":               {"POST": {`/images/create`}},
	"IMAGES_INSPECT":              {"GET": {`/images/\S+/json`}},
	"IMAGES_HISTORY":              {"GET": {`/images/\S+/history`}},
	"IMAGES_PUSH":                 {"POST": {`/images/\S+/push`}},
	"IMAGES_TAG":                  {"POST": {`/images/\S+/tag`}},
	"IMAGES_DELETE":               {"DELETE": {`/images/\S+`}},
	"IMAGES_SEARCH":               {"GET": {`/images/search`}},
	"IMAGES_PRUNE":                {"POST": {`/images/prune`}},
	"IMAGES_COMMIT":               {"POST": {`/commit`}},
	"IMAGES_GET":                  {"GET": {`/images/\S+/get`}},
	"IMAGES_GET_ALL":              {"GET": {`/images/get`}},
	"IMAGES_LOAD":                 {"POST": {`/images/load`}},
	"CONTAINERS_EXEC":             {"POST": {`/containers/\S+/exec`}},
	"EXEC_START":                  {"POST": {`/exec/\S+/start`}},
	"EXEC_RESIZE":                 {"POST": {`/exec/\S+/resize`}},
	"EXEC_INSPECT":                {"GET": {`/exec/\S+/json`}},
	"VOLUMES_LIST":                {"GET": {`/volumes`}},
	"VOLUMES_CREATE":              {"POST": {`/volumes/create`}},
	"VOLUMES_INSPECT":             {"GET": {`/volumes/\S+`}},
	"VOLUMES_UPDATE":              {"PUT": {`/volumes/\S+`}},
	"VOLUMES_DELETE":              {"DELETE": {`/volumes/\S+`}},
	"VOLUMES_PRUNE":               {"POST": {`/volumes/prune`}},
	"NETWORKS_LIST":               {"GET": {`/networks`}},
	"NETWORKS_INSPECT":            {"GET": {`/networks/\S+`}},
	"NETWORKS_DELETE":             {"DELETE": {`/networks/\S+`}},
	"NETWORKS_CREATE":             {"POST": {`/networks/create`}},
	"NETWORKS_CONNECT":            {"POST": {`/networks/\S+/connect`}},
	"NETWORKS_DISCONNECT":         {"POST": {`/networks/\S+/disconnect`}},
	"NETWORKS_PRUNE":              {"POST": {`/networks/prune`}},
	"PLUGINS_LIST":                {"GET": {`/plugins`}},
	"PLUGINS_PRIVILEGES":          {"GET": {`/plugins/privileges`}},
	"PLUGINS_PULL":                {"POST": {`/plugins/pull`}},
	"PLUGINS_INSPECT":             {"GET": {`/plugins/\S+/json`}},
	"PLUGINS_DELETE":              {"DELETE": {`/plugins/\S+`}},
	"PLUGINS_ENABLE":              {"POST": {`/plugins/\S+/enable`}},
	"PLUGINS_DISABLE":             {"POST": {`/plugins/\S+/disable`}},
	"PLUGINS_UPGRADE":             {"POST": {`/plugins/\S+/upgrade`}},
	"PLUGINS_CREATE":              {"POST": {`/plugins/create`}},
	"PLUGINS_PUSH":                {"POST": {`/plugins/\S+/push`}},
	"PLUGINS_SET":                 {"POST": {`/plugins/\S+/set`}},
	"NODES_LIST":                  {"GET": {`/nodes`}},
	"NODES_INSPECT":               {"GET": {`/nodes/\S+`}},
	"NODES_DELETE":                {"DELETE": {`/nodes/\S+`}},
	"NODES_UPDATE":                {"POST": {`/nodes/\S+/update`}},
	"SWARM_INIT":                  {"POST": {`/swarm/init`}},
	"SWARM_JOIN":                  {"POST": {`/swarm/join`}},
	"SWARM_LEAVE":                 {"POST": {`/swarm/leave`}},
	"SWARM_UPDATE":                {"POST": {`/swarm/update`}},
	"SWARM_UNLOCKKEY":             {"GET": {`/swarm/unlockkey`}},
	"SWARM_UNLOCK":                {"POST": {`/swarm/unlock`}},
	"SERVICES_LIST":               {"GET": {`/services`}},
	"SERVICES_CREATE":             {"POST": {`/services/create`}},
	"SERVICES_INSPECT":            {"GET": {`/services/\S+`}},
	"SERVICES_DELETE":             {"DELETE": {`/services/\S+`}},
	"SERVICES_UPDATE":             {"POST": {`/services/\S+/update`}},
	"SERVICES_LOGS":               {"GET": {`/services/\S+/logs`}},
	"TASKS_LIST":                  {"GET": {`/tasks`}},
	"TASKS_INSPECT":               {"GET": {`/tasks/\S+`}},
	"TASKS_LOGS":                  {"GET": {`/tasks/\S+/logs`}},
	"SECRETS_LIST":                {"GET": {`/secrets`}},
	"SECRETS_CREATE":              {"POST": {`/secrets/create`}},
	"SECRETS_INSPECT":             {"GET": {`/secrets/\S+`}},
	"SECRETS_DELETE":              {"DELETE": {`/secrets/\S+`}},
	"SECRETS_UPDATE":              {"POST": {`/secrets/\S+/update`}},
	"CONFIGS_LIST":                {"GET": {`/configs`}},
	"CONFIGS_CREATE":              {"POST": {`/configs/create`}},
	"CONFIGS_INSPECT":             {"GET": {`/configs/\S+`}},
	"CONFIGS_DELETE":              {"DELETE": {`/configs/\S+`}},
	"CONFIGS_UPDATE":              {"POST": {`/configs/\S+/update`}},
}
