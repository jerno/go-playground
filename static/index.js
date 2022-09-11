import { createApp } from "vue";
import App from './App.js'
import Topbar from './Topbar.js'

const app = createApp(App);

app.component('top-bar', Topbar);

app.mount("#app");
