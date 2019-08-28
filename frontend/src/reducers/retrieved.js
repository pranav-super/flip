import { retrieve } from '../data';

const actions = {
  updateList: 'UPDATE_RETRIEVED_LIST'
};

const initial = [];

export const updateRetrieved = id => (dispatch) => {
  return new Promise((res, rej) => {
    retrieve(id).then(r => {
        dispatch({
          type: actions.updateList,
          list: r,
        });
        res(r);
      } // TODO: Rejection
    );
  });
};

const retrieved = (state = initial, action) => {
  if (action.type === actions.updateList) {
    return action.list
  }

  return state
};

export default retrieved;