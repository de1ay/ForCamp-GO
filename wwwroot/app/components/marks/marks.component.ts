import {Component, OnInit} from '@angular/core';
import {CookieService} from 'angular2-cookie/core';
import {CheckTokenService} from '../../src/checkToken.service';
import {UserService} from '../../src/user.service';
import {OrgSetService} from "../../src/orgset.service";
import {alert} from "notie";
import {MarksService} from "../../src/marks.service";

interface Mark{
    id: number
    value: number
}

interface Participant{
    login: string
    name: string
    surname: string
    middlename: string
    sex: number
    team: number
    marks: Mark[]
}

@Component({
    selector: "marks",
    templateUrl: "app/components/marks/marks.component.html",
    styleUrls: ["app/components/marks/marks.component.css"]
})
export class MarksComponent implements OnInit {
    public Login_Edit: string = "";
    public CategoryID_Edit: number = 0;
    private Token: string;

    constructor(private cookieService: CookieService,
                private checkTokenService: CheckTokenService,
                public userService: UserService,
                public orgSetService: OrgSetService,
                public marksService: MarksService) {
    }

    ngOnInit() {
        this.TokenInit();
        this.ServiceInit();
    }

    public EditMark(login: string, categoryID: number){
        if(this.CheckPermissions(categoryID)){
            if (this.CheckSelfTeamMarks(login)) {
                this.Login_Edit = login;
                this.CategoryID_Edit = categoryID;
                this.marksService.MarkEdit_Active = true;
            } else {
                alert({type: 3, text: "Вы не можете изменять баллы своей команде", time: 2});
            }
        } else {
            alert({type: 3, text: "Вы не можете редактировать данную категорию", time: 2});
        }
    }

    private ServiceInit(){
        this.UserServiceInit();
        this.OrgSetServiceInit();
    }

    private OrgSetServiceInit(){
        if(this.orgSetService.Token == undefined){
            this.orgSetService.Token = this.Token;
        }
        this.orgSetService.GetOrgSettings();
        this.orgSetService.GetCategories();
        this.orgSetService.GetTeams();
        this.orgSetService.GetParticipants();
        this.orgSetService.GetEmployees();
    }

    private UserServiceInit(){
        if(this.userService.Token == undefined) {
            this.userService.Token = this.Token;
        }
        this.userService.GetSelfUserData();
    }

    private CheckSelfTeamMarks(login: string): boolean{
        if (!this.orgSetService.OrgSettings.self_marks){
            if (this.userService.SelfData.Team != 0){
                for (let i = 0; i < this.orgSetService.Teams.length; i++){
                    if (this.orgSetService.Teams[i].id == this.userService.SelfData.Team){
                        for (let k = 0; k < this.orgSetService.Teams[i].participants.length; k++){
                            if (this.orgSetService.Teams[i].participants[k] == login){
                                return false
                            }
                        }
                        return true
                    }
                }
            } else {
                return true
            }
        } else {
            return true
        }
    }

    private CheckPermissions(categoryID: number): boolean{
        for(let i = 0; i < this.userService.SelfData.Permissions.length; i++){
            if (this.userService.SelfData.Permissions[i].id == categoryID){
                return this.userService.SelfData.Permissions[i].value;
            }
        }
        return false
    }

    private TokenInit(){
        this.Token = this.cookieService.get("token");
        if (this.Token == undefined) {
            window.location.href = "https://forcamp.ga/exit";
        }
        this.checkTokenService.CheckToken(this.Token);
    }

}