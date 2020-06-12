<template>
  <client-only>
    <mavon-editor v-model="article.Context"
                  defaultOpen="preview"
                  :subfield="false"
                  :toolbarsFlag="false"
                  :navigation="true"
                  :boxShadow="true"
                  :ishljs="true" />
  </client-only>
</template>

<script>
import { getArticleByID } from "@/api/api"
import { Message } from 'element-ui'
export default {
  data () {
    return {
      article: {},
    }
  },

  asyncData ({ query }) {
    return getArticleByID(query).then(res => {
      if (res.data.state === true) {
        return { article: res.data.data }
      } else {
        Message({
          message: res.data.msg,
          type: 'error',
          duration: 3 * 1000
        })
      }
    }).catch(err => console.log(err));
  },
}
</script>