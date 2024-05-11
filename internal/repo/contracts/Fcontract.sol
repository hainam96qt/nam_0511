// SPDX-License-Identifier: MIT

pragma solidity ^0.8.19;

contract AttendanceContract {
    struct AttendanceRecord {
        uint employeeId;
        uint checkInTime;
        string details;
    }

    uint256 private number = 1;

    mapping(uint => AttendanceRecord[]) private attendanceRecords;
    mapping(address => uint256) private authorizedEntities;

    event AttendanceRecorded(uint indexed employeeId, uint checkInTime, string details);
    event AttendanceUpdated(uint indexed employeeId, uint index, uint newCheckInTime, string newDetails);

    constructor() {
        authorizedEntities[msg.sender] = 1000000; // Deployer is authorized by default
    }

    // user ton tai
    modifier onlyAuthorized() {
        require(authorizedEntities[msg.sender] != 0, "Unauthorized access");
        _;
    }

    // truong hop la chinh chu user hoac admin thi dc tao
    function recordAttendance(uint _employeeId, uint _checkInTime, string memory _details) public onlyAuthorized {
        if (authorizedEntities[msg.sender] != 1000000  && _employeeId != authorizedEntities[msg.sender]){
            return;
        }
        attendanceRecords[_employeeId].push(AttendanceRecord(_employeeId, _checkInTime, _details));
        emit AttendanceRecorded(_employeeId, _checkInTime, _details);
    }

    function getAttendanceByEmployeeId(uint _employeeId) public view returns (AttendanceRecord[] memory) {
        return attendanceRecords[_employeeId];
    }

    function getAttendanceByDateRange(uint _employeeId, uint _start, uint _end) public view returns (AttendanceRecord[] memory) {
        AttendanceRecord[] memory records = attendanceRecords[_employeeId];
        AttendanceRecord[] memory filteredRecords = new AttendanceRecord[](records.length);
        for (uint i = 0; i < records.length; i++) {
            if (records[i].checkInTime >= _start && records[i].checkInTime <= _end) {
                filteredRecords[i]=records[i];
            }
        }
        return filteredRecords;
    }

    function updateAttendance(uint _employeeId, uint _index, uint _newCheckInTime, string memory _newDetails) public onlyAuthorized {
        require(_index < attendanceRecords[_employeeId].length, "Index out of bounds");
        attendanceRecords[_employeeId][_index].checkInTime = _newCheckInTime;
        attendanceRecords[_employeeId][_index].details = _newDetails;
        emit AttendanceUpdated(_employeeId, _index, _newCheckInTime, _newDetails);
    }

    function getEmployeeID() public onlyAuthorized view returns (uint _emID){
        require(authorizedEntities[msg.sender] != 0, "Unauthorized access");

        return authorizedEntities[msg.sender];
    }

    function authorizeEntity(address _entity) public onlyAuthorized {
        number = number++;
        authorizedEntities[_entity] = number;
    }

    function revokeAuthorization(address _entity) public onlyAuthorized {
        authorizedEntities[_entity] = 0;
    }
}