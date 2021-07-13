use cosmwasm_std::{
    to_binary, Deps, DepsMut, Env, HandleResponse, InitResponse, MessageInfo,Binary,StdResult
};

use crate::error::ContractError;
use crate::msg::{HandleMsg, InitMsg,QueryMsg};
use crate::state::{State,config};

// Note, you can use StdResult in some functions where you do not
// make use of the custom errors
pub fn init(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    _msg: InitMsg,
) -> Result<InitResponse, ContractError> {
    let state = State {
        owner: deps.api.canonical_address(&info.sender)?,
    };
    config(deps.storage).save(&state)?;

    Ok(InitResponse::default())
}

// And declare a custom Error variant for the ones where you will want to make use of it
pub fn handle(
    deps: DepsMut,
    _env: Env,
    _info: MessageInfo,
    msg: HandleMsg,
) -> Result<HandleResponse, ContractError> {
    match msg {
        HandleMsg::Hello{words} => try_hello(deps,words),
    }
}

pub fn try_hello(_deps: DepsMut, words: String) -> Result<HandleResponse, ContractError> {
    Ok(HandleResponse {
        messages: vec![],
        data: Some(Binary::from(words.as_bytes())),
        attributes: vec![],
    })
}

pub fn query(_deps: Deps, _env: Env, _msg: QueryMsg) -> StdResult<Binary> {
    to_binary("no query function")
}