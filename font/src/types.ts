import type { Dayjs } from 'dayjs';

export interface UserFormData {
    FNameValue: string;
    LNameValue: string;
    IDValue: string;
    UserNameValue: string;
    PasswordValue: string;
    Password2Value: string;
    EmailValue: string;
    PhoneNumberValue: string;
}

export interface CompanyFormData {
    CompanyNameValue: string;
    CompanyIDValue: string;
    AddressValue: string;
}

export interface MessageResponse {
    message: string;
}

export interface DriverHeaderData {
    AddressValue: string;
    IsEmployeeValue: boolean;
}

export interface DriverListItem {
    nameOfDriver: string;
    aviable: boolean;
}

export interface GetDriversResponse {
    message: string;
    List: DriverListItem[];
}

export interface DriveListItem {
    nameOfDrive: string;
    aviable: boolean;
}

export interface GetDrivesResponse {
    message: string;
    List: DriveListItem[];
}

export interface DriveFormData {
    NameValue: string;
    SelectTimeValue: Dayjs | null;
    StartGpsValue: string;
    EndGpsValue: string;
}
