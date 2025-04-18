const manifest = (() => {
function __memo(fn) {
	let value;
	return () => value ??= (value = fn());
}

return {
	appDir: "_app",
	appPath: "_app",
	assets: new Set(["favicon.png"]),
	mimeTypes: {".png":"image/png"},
	_: {
		client: {start:"_app/immutable/entry/start.fyieEP8T.js",app:"_app/immutable/entry/app.Rg_twRHy.js",imports:["_app/immutable/entry/start.fyieEP8T.js","_app/immutable/chunks/Cv5cQigG.js","_app/immutable/chunks/Br5c-oLs.js","_app/immutable/chunks/Cgeft74Y.js","_app/immutable/chunks/DIS1lw11.js","_app/immutable/entry/app.Rg_twRHy.js","_app/immutable/chunks/Cgeft74Y.js","_app/immutable/chunks/5z2gxC1j.js","_app/immutable/chunks/DJoc1YTT.js","_app/immutable/chunks/CoWB_c7W.js","_app/immutable/chunks/tsR4g0rO.js","_app/immutable/chunks/DIS1lw11.js","_app/immutable/chunks/BqR-gka9.js","_app/immutable/chunks/DOLpMf4M.js","_app/immutable/chunks/Br5c-oLs.js"],stylesheets:[],fonts:[],uses_env_dynamic_public:false},
		nodes: [
			__memo(() => import('./chunks/0-D1oa7eZU.js')),
			__memo(() => import('./chunks/1-KNKE-LFK.js')),
			__memo(() => import('./chunks/2-BC7EXo7t.js')),
			__memo(() => import('./chunks/3-BnnvZngq.js')),
			__memo(() => import('./chunks/4-N5QRbIY0.js').then(function (n) { return n._; })),
			__memo(() => import('./chunks/5-biIfybmS.js').then(function (n) { return n._; })),
			__memo(() => import('./chunks/6-bWLdTWOK.js')),
			__memo(() => import('./chunks/7-Dlqh8Rsr.js').then(function (n) { return n._; })),
			__memo(() => import('./chunks/8-DZZofe2q.js'))
		],
		routes: [
			{
				id: "/",
				pattern: /^\/$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 2 },
				endpoint: null
			},
			{
				id: "/pages/about",
				pattern: /^\/pages\/about\/?$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 3 },
				endpoint: null
			},
			{
				id: "/pages/admin-page",
				pattern: /^\/pages\/admin-page\/?$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 4 },
				endpoint: null
			},
			{
				id: "/pages/event-room",
				pattern: /^\/pages\/event-room\/?$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 5 },
				endpoint: null
			},
			{
				id: "/pages/event-room/challenge",
				pattern: /^\/pages\/event-room\/challenge\/?$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 6 },
				endpoint: null
			},
			{
				id: "/pages/new-event",
				pattern: /^\/pages\/new-event\/?$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 7 },
				endpoint: null
			},
			{
				id: "/pages/signin",
				pattern: /^\/pages\/signin\/?$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 8 },
				endpoint: null
			}
		],
		prerendered_routes: new Set([]),
		matchers: async () => {
			
			return {  };
		},
		server_assets: {}
	}
}
})();

const prerendered = new Set([]);

const base = "";

export { base, manifest, prerendered };
//# sourceMappingURL=manifest.js.map
