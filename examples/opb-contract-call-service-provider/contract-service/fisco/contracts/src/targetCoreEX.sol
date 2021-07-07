pragma solidity ^0.4.24;

/**
 * @title iService Core Extension contract
 */
contract targetCoreEx {
    mapping(bytes32 => bool) requests;

    /**
    * @dev Event triggered when the request is sent
    * @param _RequestID Request id
    * @param _result result bytes
    */
    event CrossChainResponseSent(
        bytes32 _RequestID,
        bytes _result
    );
    /**
    * @dev Make sure that the Request  has not been responded
    * @param _RequestID Request id
    */
    modifier validateRequest(bytes32 _RequestID) {
        require(
            requests[_RequestID] == false,
            "iServiceCoreEx: duplicated request!"
        );

        _;
    }

    /**
     * @dev call service/contract in dest chain
     * @param _RequestID Request id
     * @param _endpointAddress endpointAddress
     * @param _callData call data from source chain
     */
    function callService(
        bytes32 _RequestID,
        address _endpointAddress,
        bytes _callData
    ) public validateRequest(_RequestID) {
        uint callDataLength = _callData.length;
        bytes memory result;
        uint success;
        assembly {
        // call
            let d := add(_callData, 0x20)
            success := call(
            gas(),
            _endpointAddress,
            callvalue,
            d,
            callDataLength,
            0,
            0
            )

        // handle result
            switch success
            case 1 {
                let size := returndatasize
                result := mload(0x40)
                mstore(0x40, add(result, and(add(add(size, 0x20), 0x1f), not(0x1f))))
                mstore(result, size)
                returndatacopy(add(result, 0x20), 0, size)
            }
            case 0 {
            // call failed and revert
                revert(0, 0)
            }
        }
        if (success == 1) {
            emit CrossChainResponseSent(
                _RequestID,
                result
            );
        }
    }
}
