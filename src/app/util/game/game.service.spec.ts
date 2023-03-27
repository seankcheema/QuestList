import { HttpClient, HttpErrorResponse } from "@angular/common/http";
import { of } from "rxjs";

import { GameService } from "./game.service";

describe("GameService", () => {
    let service: GameService;
    let httpClientSpy: jasmine.SpyObj<HttpClient>;
    
    beforeEach(() => {
        httpClientSpy = jasmine.createSpyObj('HttpClient', ['get']);
        service = new GameService(httpClientSpy);
    });
    
    it("should be created", () => {
        expect(service).toBeTruthy();
    });

    it('#getGames() should return value from observable', (done: DoneFn) => {

        const expectedGameSlug:string = 'grand-theft-auto-v';

        httpClientSpy.get.and.returnValue(of([{
            slug: expectedGameSlug
        }]));

        service.getGames('1').subscribe(value => {
            expect(value[0].slug)
                .withContext('expected game slug')
                .toEqual(expectedGameSlug);
            done();
        });
        
        expect(httpClientSpy.get.calls.count())
            .withContext('one call')
            .toBe(1);
    });

    it('#getRecentGames() should return value from observable', (done: DoneFn) => {
            
            const expectedGameSlug:string = 'most recent game name';
    
            httpClientSpy.get.and.returnValue(of([{
                slug: expectedGameSlug
            }]));
    
            service.getRecentGames().subscribe(value => {
                expect(value[0].slug)
                    .withContext('expected game slug')
                    .toEqual(expectedGameSlug);
                done();
            });
            
            expect(httpClientSpy.get.calls.count())
                .withContext('one call')
                .toBe(1);
    });
});