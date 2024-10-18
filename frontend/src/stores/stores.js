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
    handleQueryForm(queryForm) {
      this.queryForm=queryForm
    },
    handleLoadingTable(value){
        this.loadingTable=value
    },
    handleTabelDatas(value){
        this.tabelDatas=value
    },
    handleTabelInfo(value){
        this.tableInfo=value
    }
  },
})

export const useThemeStore = defineStore('theme', {
  state: () => ({
    main:{},
    mode:'',
  }),
  actions: {
    handleTheme(theme) {
      this.main=theme
    },
    handleMode(mode) {
      this.mode=mode
    },
  },
})

export const appInfoStore = defineStore('appInfo', {
  state: () => ({
    version: 'v0.0.1 Preview',
    main:'Axolotland',
    log:'',
  }),
})