import Vue from 'vue';
import Vuetify from 'vuetify/lib/framework';

Vue.use(Vuetify);

export default new Vuetify({
    theme: {
        themes : {
            light: {
                secondary : "#FFD369",
                background : '#222831',
                background2 : '#F4F9F9',
            }
        },
        options: {
            customProperties: true
        }
    }
});
