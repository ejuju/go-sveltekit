import adapter from '@sveltejs/adapter-static';
import preprocess from 'svelte-preprocess';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	preprocess: preprocess(),
	kit: {
		appDir: 'appdir',
		adapter: adapter(),
		trailingSlash: 'always'
	}
};

export default config;
