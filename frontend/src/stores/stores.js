import { defineStore } from 'pinia'

export const stateStore = defineStore('state', {
  state: () => ({
    queryForm: {},
    loadingTable:false,
    tableInfo:{
        current:1,
        pageSize:10,
    },
    tabelDatas:{
        count:0,
        field:[],
        data:[],
    },
  }),
  actions: {
    handelQueryForm(queryForm) {
      this.queryForm=queryForm
    },
    handelLoadingTable(value){
        this.loadingTable=value
    },
    handelTabelDatas(value){
        this.tabelDatas=value
    },
    handelTabelInfo(value){
        this.tableInfo=value
    }
  },
})

export const appInfoStore = defineStore('appInfo', {
  state: () => ({
    version: 'v0.0.1 Preview',
    log:'',
  }),
})