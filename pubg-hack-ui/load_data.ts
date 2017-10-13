interface EvilLoadData {
  real_load_js: string;
  api_host: string;
}

declare var LOAD_DATA: EvilLoadData;

export const evilLoadData = LOAD_DATA || {real_load_js: 'console.log("no load JS!")', api_host: 'nil'};
