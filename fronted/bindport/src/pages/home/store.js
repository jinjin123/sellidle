import { observable } from "mobx";
import http from "../../libs/http";

class Store {
  @observable records = [];
  @observable isFetching = false;
  @observable PortForm = false;

  fetchRecords = () => {
    this.isFetching = true;
    return http.get('/api/check')
      .then((res) => {
        this.records = res
      })
      .finally(() => this.isFetching = false)

  }
  showBindPortForm = () => {
    this.PortForm = true;
  }
}


export default new Store()

