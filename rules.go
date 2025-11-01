package main

type rule struct {
	enabled bool
	description string
	routes map[string][]string
}

var rules = map[string]*rule{
	"allow-container-list": {
		description: "Allow GET /containers/json",
		routes: map[string][]string{ "GET": {"/containers/json"} },
	},
	"allow-container-create": {
		description: "Allow POST /containers/create",
		routes: map[string][]string{ "POST": {"/containers/create"} },
	},
	"allow-container-inspect": {
		description: "Allow GET /containers/{id}/json",
		routes: map[string][]string{ "GET": {"/containers/*/json"} },
	},
	"allow-container-top": {
		description: "Allow GET /containers/{id}/top",
		routes: map[string][]string{ "GET": {"/containers/*/top"} },
	},
	"allow-container-logs": {
		description: "Allow GET /containers/{id}/logs",
		routes: map[string][]string{ "GET": {"/containers/*/logs"} },
	},
	"allow-container-changes": {
		description: "Allow GET /containers/{id}/changes",
		routes: map[string][]string{ "GET": {"/containers/*/changes"} },
	},
	"allow-container-export": {
		description: "Allow GET /containers/{id}/export",
		routes: map[string][]string{ "GET": {"/containers/*/export"} },
	},
	"allow-container-stats": {
		description: "Allow GET /containers/{id}/stats",
		routes: map[string][]string{ "GET": {"/containers/*/stats"} },
	},
	"allow-container-resize": {
		description: "Allow POST /containers/{id}/resize",
		routes: map[string][]string{ "POST": {"/containers/*/resize"} },
	},
	"allow-container-start": {
		description: "Allow POST /containers/{id}/start",
		routes: map[string][]string{ "POST": {"/containers/*/start"} },
	},
	"allow-container-stop": {
		description: "Allow POST /containers/{id}/stop",
		routes: map[string][]string{ "POST": {"/containers/*/stop"} },
	},
	"allow-container-restart": {
		description: "Allow POST /containers/{id}/restart",
		routes: map[string][]string{ "POST": {"/containers/*/restart"} },
	},
	"allow-container-kill": {
		description: "Allow POST /containers/{id}/kill",
		routes: map[string][]string{ "POST": {"/containers/*/kill"} },
	},
	"allow-container-update": {
		description: "Allow POST /containers/{id}/update",
		routes: map[string][]string{ "POST": {"/containers/*/update"} },
	},
	"allow-container-rename": {
		description: "Allow POST /containers/{id}/rename",
		routes: map[string][]string{ "POST": {"/containers/*/rename"} },
	},
	"allow-container-pause": {
		description: "Allow POST /containers/{id}/pause",
		routes: map[string][]string{ "POST": {"/containers/*/pause"} },
	},
	"allow-container-unpause": {
		description: "Allow POST /containers/{id}/unpause",
		routes: map[string][]string{ "POST": {"/containers/*/unpause"} },
	},
	"allow-container-attach": {
		description: "Allow POST /containers/{id}/attach",
		routes: map[string][]string{ "POST": {"/containers/*/attach"} },
	},
	"allow-container-attach-websocket": {
		description: "Allow GET /containers/{id}/attach/ws",
		routes: map[string][]string{ "GET": {"/containers/*/attach/ws"} },
	},
	"allow-container-wait": {
		description: "Allow POST /containers/{id}/wait",
		routes: map[string][]string{ "POST": {"/containers/*/wait"} },
	},
	"allow-container-delete": {
		description: "Allow DELETE /containers/{id}",
		routes: map[string][]string{ "DELETE": {"/containers/*"} },
	},
	"allow-container-archive-info": {
		description: "Allow HEAD /containers/{id}/archive",
		routes: map[string][]string{ "HEAD": {"/containers/*/archive"} },
	},
	"allow-container-archive": {
		description: "Allow GET /containers/{id}/archive",
		routes: map[string][]string{ "GET": {"/containers/*/archive"} },
	},
	"allow-put-container-archive": {
		description: "Allow PUT /containers/{id}/archive",
		routes: map[string][]string{ "PUT": {"/containers/*/archive"} },
	},
	"allow-container-prune": {
		description: "Allow POST /containers/prune",
		routes: map[string][]string{ "POST": {"/containers/prune"} },
	},
	"allow-image-list": {
		description: "Allow GET /images/json",
		routes: map[string][]string{ "GET": {"/images/json"} },
	},
	"allow-image-build": {
		description: "Allow POST /build",
		routes: map[string][]string{ "POST": {"/build"} },
	},
	"allow-build-prune": {
		description: "Allow POST /build/prune",
		routes: map[string][]string{ "POST": {"/build/prune"} },
	},
	"allow-image-create": {
		description: "Allow POST /images/create",
		routes: map[string][]string{ "POST": {"/images/create"} },
	},
	"allow-image-inspect": {
		description: "Allow GET /images/{name}/json",
		routes: map[string][]string{ "GET": {"/images/*/json"} },
	},
	"allow-image-history": {
		description: "Allow GET /images/{name}/history",
		routes: map[string][]string{ "GET": {"/images/*/history"} },
	},
	"allow-image-push": {
		description: "Allow POST /images/{name}/push",
		routes: map[string][]string{ "POST": {"/images/*/push"} },
	},
	"allow-image-tag": {
		description: "Allow POST /images/{name}/tag",
		routes: map[string][]string{ "POST": {"/images/*/tag"} },
	},
	"allow-image-delete": {
		description: "Allow DELETE /images/{name}",
		routes: map[string][]string{ "DELETE": {"/images/*"} },
	},
	"allow-image-search": {
		description: "Allow GET /images/search",
		routes: map[string][]string{ "GET": {"/images/search"} },
	},
	"allow-image-prune": {
		description: "Allow POST /images/prune",
		routes: map[string][]string{ "POST": {"/images/prune"} },
	},
	"allow-system-auth": {
		description: "Allow POST /auth",
		routes: map[string][]string{ "POST": {"/auth"} },
	},
	"allow-system-info": {
		description: "Allow GET /info",
		routes: map[string][]string{ "GET": {"/info"} },
	},
	"allow-system-version": {
		description: "Allow GET /version",
		routes: map[string][]string{ "GET": {"/version"} },
	},
	"allow-system-ping": {
		description: "Allow GET /_ping",
		routes: map[string][]string{ "GET": {"/_ping"} },
	},
	"allow-system-ping-head": {
		description: "Allow HEAD /_ping",
		routes: map[string][]string{ "HEAD": {"/_ping"} },
	},
	"allow-image-commit": {
		description: "Allow POST /commit",
		routes: map[string][]string{ "POST": {"/commit"} },
	},
	"allow-system-events": {
		description: "Allow GET /events",
		routes: map[string][]string{ "GET": {"/events"} },
	},
	"allow-system-data-usage": {
		description: "Allow GET /system/df",
		routes: map[string][]string{ "GET": {"/system/df"} },
	},
	"allow-image-get": {
		description: "Allow GET /images/{name}/get",
		routes: map[string][]string{ "GET": {"/images/*/get"} },
	},
	"allow-image-get-all": {
		description: "Allow GET /images/get",
		routes: map[string][]string{ "GET": {"/images/get"} },
	},
	"allow-image-load": {
		description: "Allow POST /images/load",
		routes: map[string][]string{ "POST": {"/images/load"} },
	},
	"allow-container-exec": {
		description: "Allow POST /containers/{id}/exec",
		routes: map[string][]string{ "POST": {"/containers/*/exec"} },
	},
	"allow-exec-start": {
		description: "Allow POST /exec/{id}/start",
		routes: map[string][]string{ "POST": {"/exec/*/start"} },
	},
	"allow-exec-resize": {
		description: "Allow POST /exec/{id}/resize",
		routes: map[string][]string{ "POST": {"/exec/*/resize"} },
	},
	"allow-exec-inspect": {
		description: "Allow GET /exec/{id}/json",
		routes: map[string][]string{ "GET": {"/exec/*/json"} },
	},
	"allow-volume-list": {
		description: "Allow GET /volumes",
		routes: map[string][]string{ "GET": {"/volumes"} },
	},
	"allow-volume-create": {
		description: "Allow POST /volumes/create",
		routes: map[string][]string{ "POST": {"/volumes/create"} },
	},
	"allow-volume-inspect": {
		description: "Allow GET /volumes/{name}",
		routes: map[string][]string{ "GET": {"/volumes/*"} },
	},
	"allow-volume-update": {
		description: "Allow PUT /volumes/{name}",
		routes: map[string][]string{ "PUT": {"/volumes/*"} },
	},
	"allow-volume-delete": {
		description: "Allow DELETE /volumes/{name}",
		routes: map[string][]string{ "DELETE": {"/volumes/*"} },
	},
	"allow-volume-prune": {
		description: "Allow POST /volumes/prune",
		routes: map[string][]string{ "POST": {"/volumes/prune"} },
	},
	"allow-network-list": {
		description: "Allow GET /networks",
		routes: map[string][]string{ "GET": {"/networks"} },
	},
	"allow-network-inspect": {
		description: "Allow GET /networks/{id}",
		routes: map[string][]string{ "GET": {"/networks/*"} },
	},
	"allow-network-delete": {
		description: "Allow DELETE /networks/{id}",
		routes: map[string][]string{ "DELETE": {"/networks/*"} },
	},
	"allow-network-create": {
		description: "Allow POST /networks/create",
		routes: map[string][]string{ "POST": {"/networks/create"} },
	},
	"allow-network-connect": {
		description: "Allow POST /networks/{id}/connect",
		routes: map[string][]string{ "POST": {"/networks/*/connect"} },
	},
	"allow-network-disconnect": {
		description: "Allow POST /networks/{id}/disconnect",
		routes: map[string][]string{ "POST": {"/networks/*/disconnect"} },
	},
	"allow-network-prune": {
		description: "Allow POST /networks/prune",
		routes: map[string][]string{ "POST": {"/networks/prune"} },
	},
	"allow-plugin-list": {
		description: "Allow GET /plugins",
		routes: map[string][]string{ "GET": {"/plugins"} },
	},
	"allow-get-plugin-privileges": {
		description: "Allow GET /plugins/privileges",
		routes: map[string][]string{ "GET": {"/plugins/privileges"} },
	},
	"allow-plugin-pull": {
		description: "Allow POST /plugins/pull",
		routes: map[string][]string{ "POST": {"/plugins/pull"} },
	},
	"allow-plugin-inspect": {
		description: "Allow GET /plugins/{name}/json",
		routes: map[string][]string{ "GET": {"/plugins/*/json"} },
	},
	"allow-plugin-delete": {
		description: "Allow DELETE /plugins/{name}",
		routes: map[string][]string{ "DELETE": {"/plugins/*"} },
	},
	"allow-plugin-enable": {
		description: "Allow POST /plugins/{name}/enable",
		routes: map[string][]string{ "POST": {"/plugins/*/enable"} },
	},
	"allow-plugin-disable": {
		description: "Allow POST /plugins/{name}/disable",
		routes: map[string][]string{ "POST": {"/plugins/*/disable"} },
	},
	"allow-plugin-upgrade": {
		description: "Allow POST /plugins/{name}/upgrade",
		routes: map[string][]string{ "POST": {"/plugins/*/upgrade"} },
	},
	"allow-plugin-create": {
		description: "Allow POST /plugins/create",
		routes: map[string][]string{ "POST": {"/plugins/create"} },
	},
	"allow-plugin-push": {
		description: "Allow POST /plugins/{name}/push",
		routes: map[string][]string{ "POST": {"/plugins/*/push"} },
	},
	"allow-plugin-set": {
		description: "Allow POST /plugins/{name}/set",
		routes: map[string][]string{ "POST": {"/plugins/*/set"} },
	},
	"allow-node-list": {
		description: "Allow GET /nodes",
		routes: map[string][]string{ "GET": {"/nodes"} },
	},
	"allow-node-inspect": {
		description: "Allow GET /nodes/{id}",
		routes: map[string][]string{ "GET": {"/nodes/*"} },
	},
	"allow-node-delete": {
		description: "Allow DELETE /nodes/{id}",
		routes: map[string][]string{ "DELETE": {"/nodes/*"} },
	},
	"allow-node-update": {
		description: "Allow POST /nodes/{id}/update",
		routes: map[string][]string{ "POST": {"/nodes/*/update"} },
	},
	"allow-swarm-inspect": {
		description: "Allow GET /swarm",
		routes: map[string][]string{ "GET": {"/swarm"} },
	},
	"allow-swarm-init": {
		description: "Allow POST /swarm/init",
		routes: map[string][]string{ "POST": {"/swarm/init"} },
	},
	"allow-swarm-join": {
		description: "Allow POST /swarm/join",
		routes: map[string][]string{ "POST": {"/swarm/join"} },
	},
	"allow-swarm-leave": {
		description: "Allow POST /swarm/leave",
		routes: map[string][]string{ "POST": {"/swarm/leave"} },
	},
	"allow-swarm-update": {
		description: "Allow POST /swarm/update",
		routes: map[string][]string{ "POST": {"/swarm/update"} },
	},
	"allow-swarm-unlockkey": {
		description: "Allow GET /swarm/unlockkey",
		routes: map[string][]string{ "GET": {"/swarm/unlockkey"} },
	},
	"allow-swarm-unlock": {
		description: "Allow POST /swarm/unlock",
		routes: map[string][]string{ "POST": {"/swarm/unlock"} },
	},
	"allow-service-list": {
		description: "Allow GET /services",
		routes: map[string][]string{ "GET": {"/services"} },
	},
	"allow-service-create": {
		description: "Allow POST /services/create",
		routes: map[string][]string{ "POST": {"/services/create"} },
	},
	"allow-service-inspect": {
		description: "Allow GET /services/{id}",
		routes: map[string][]string{ "GET": {"/services/*"} },
	},
	"allow-service-delete": {
		description: "Allow DELETE /services/{id}",
		routes: map[string][]string{ "DELETE": {"/services/*"} },
	},
	"allow-service-update": {
		description: "Allow POST /services/{id}/update",
		routes: map[string][]string{ "POST": {"/services/*/update"} },
	},
	"allow-service-logs": {
		description: "Allow GET /services/{id}/logs",
		routes: map[string][]string{ "GET": {"/services/*/logs"} },
	},
	"allow-task-list": {
		description: "Allow GET /tasks",
		routes: map[string][]string{ "GET": {"/tasks"} },
	},
	"allow-task-inspect": {
		description: "Allow GET /tasks/{id}",
		routes: map[string][]string{ "GET": {"/tasks/*"} },
	},
	"allow-task-logs": {
		description: "Allow GET /tasks/{id}/logs",
		routes: map[string][]string{ "GET": {"/tasks/*/logs"} },
	},
	"allow-secret-list": {
		description: "Allow GET /secrets",
		routes: map[string][]string{ "GET": {"/secrets"} },
	},
	"allow-secret-create": {
		description: "Allow POST /secrets/create",
		routes: map[string][]string{ "POST": {"/secrets/create"} },
	},
	"allow-secret-inspect": {
		description: "Allow GET /secrets/{id}",
		routes: map[string][]string{ "GET": {"/secrets/*"} },
	},
	"allow-secret-delete": {
		description: "Allow DELETE /secrets/{id}",
		routes: map[string][]string{ "DELETE": {"/secrets/*"} },
	},
	"allow-secret-update": {
		description: "Allow POST /secrets/{id}/update",
		routes: map[string][]string{ "POST": {"/secrets/*/update"} },
	},
	"allow-config-list": {
		description: "Allow GET /configs",
		routes: map[string][]string{ "GET": {"/configs"} },
	},
	"allow-config-create": {
		description: "Allow POST /configs/create",
		routes: map[string][]string{ "POST": {"/configs/create"} },
	},
	"allow-config-inspect": {
		description: "Allow GET /configs/{id}",
		routes: map[string][]string{ "GET": {"/configs/*"} },
	},
	"allow-config-delete": {
		description: "Allow DELETE /configs/{id}",
		routes: map[string][]string{ "DELETE": {"/configs/*"} },
	},
	"allow-config-update": {
		description: "Allow POST /configs/{id}/update",
		routes: map[string][]string{ "POST": {"/configs/*/update"} },
	},
	"allow-distribution-inspect": {
		description: "Allow GET /distribution/{name}/json",
		routes: map[string][]string{ "GET": {"/distribution/*/json"} },
	},
	"allow-session": {
		description: "Allow POST /session",
		routes: map[string][]string{ "POST": {"/session"} },
	},
}
