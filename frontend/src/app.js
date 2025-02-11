import htmx from 'htmx.org';
import Alpine from 'alpinejs';

window.htmx = htmx;
console.log('htmx is initialized');

window.Alpine = Alpine;
Alpine.start();
document.body.addEventListener('htmx:afterSwap', () => {
  Alpine.initTree(document);
});
console.log('Alpine.js is initialized');
