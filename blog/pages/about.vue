<template>
  <client-only style="box-sizing:border-box;">
    <mavon-editor v-model="article.Describe"
                  :subfield="false"
                  defaultOpen="preview"
                  :toolbarsFlag="false"
                  :boxShadow="false"
                  :ishljs="true" />
  </client-only>
</template>

<script>
import { getBloggerInfo } from "@/api/api"
import { Message } from 'element-ui'
export default {
  data () {
    return {
      article: {},
    }
  },

  asyncData () {
    return getBloggerInfo({id:1}).then(res => {
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