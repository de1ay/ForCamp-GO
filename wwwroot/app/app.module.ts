import {NgModule}      from '@angular/core';
import {BrowserModule} from '@angular/platform-browser';
import {FormsModule}   from '@angular/forms';
import {HttpModule} from '@angular/http';
import {MaterialModule} from '@angular/material';
import {RouterModule, Routes} from "@angular/router";
import {NgxDatatableModule} from "@swimlane/ngx-datatable";
import {BrowserAnimationsModule} from "@angular/platform-browser/animations";
import 'hammerjs'

import {CheckTokenService} from './src/checkToken.service';
import {UserService} from './src/user.service';
import {CookieService} from 'angular2-cookie/services/cookies.service';
import {MDL} from "./MDLInit";

import {OrgMainComponent} from "./components/orgmain/orgmain.component";
import {AuthorizationComponent} from "./components/authorization/authorization.component";
import {AppComponent} from "./app.component";
import {OrgSetComponent} from "./components/orgset/orgset.componenet";
import {OrgSetService} from "./src/orgset.service";
import {ParticipantValueEditComponent} from "./components/orgset/participantValueEdit/participantValueEdit.component";
import {PeriodValueEditComponent} from "./components/orgset/periodValueEdit/periodValueEdit.component";
import {TeamValueEditComponent} from "./components/orgset/teamValueEdit/teamValueEdit.component";
import {OrganizationValueEditComponent} from "./components/orgset/organizationValueEdit/organizationValueEdit.component";
import {AddCategoryComponent} from "./components/orgset/addCategory/addCategory.component";
import {AddTeamComponent} from "./components/orgset/addTeam/addTeam.component";
import {AddParticipantComponent} from "./components/orgset/addParticipant/addParticipant.component";
import {AddEmployeeComponent} from "./components/orgset/addEmployee/addEmployee.component";

const appRoutes: Routes = [
    {path: 'orgset', component: OrgSetComponent},
    {path: 'main', component: OrgMainComponent},
    {path: '', component: AuthorizationComponent},
];
@NgModule({
    imports: [BrowserModule, FormsModule, HttpModule, MaterialModule, RouterModule.forRoot(appRoutes), NgxDatatableModule, BrowserAnimationsModule ],
    declarations: [OrgMainComponent,
        AuthorizationComponent,
        AppComponent,
        OrgSetComponent,
        MDL,
        ParticipantValueEditComponent,
        PeriodValueEditComponent,
        TeamValueEditComponent,
        OrganizationValueEditComponent,
        AddCategoryComponent,
        AddTeamComponent,
        AddParticipantComponent,
        AddEmployeeComponent],
    bootstrap: [AppComponent],
    providers: [CookieService, CheckTokenService, UserService, OrgSetService],
})
export class AppModule {
}