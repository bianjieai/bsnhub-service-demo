use cw_storage_plus::Map;

// REQUESTS stores the requests whether the response has been executed: <request_id, responded>
pub const REQUESTS: Map<&str, bool> = Map::new("requests");