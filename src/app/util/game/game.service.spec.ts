import { HttpClient } from "@angular/common/http";
import { of } from "rxjs";

import { GameService } from "./game.service";

describe("GameService", () => {
    let service: GameService;
    let httpClientSpy: jasmine.SpyObj<HttpClient>;
    
    /**
     * Setup the service and spy before each test
     */
    beforeEach(() => {
        httpClientSpy = jasmine.createSpyObj('HttpClient', ['get']);
        service = new GameService(httpClientSpy);
    });
    
    /**
     * Test that the service is created
     */
    it("should be created", () => {
        expect(service).toBeTruthy();
    });

    /**
     * Test that the getGames method returns the expected value
     */
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

    /**
     * Test that the getRecentGames method returns the expected value
     */
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

    /*
    * Test that the getGame method returns the expected value
    */
    it('#getGame() should return value from observable', (done: DoneFn) => {

        const expectedGameSlug:string = 'grand-theft-auto-v';

        httpClientSpy.get.and.returnValue(of([{
            slug: expectedGameSlug
        }]));

        service.getGame('grand-theft-auto-v').subscribe(value => {
            expect(value[0].slug)
                .withContext('expected game slug')
                .toEqual(expectedGameSlug);
            done();
        });
        
        expect(httpClientSpy.get.calls.count())
            .withContext('one call')
            .toBe(1);
    })

    /**
     * Test that the getUpcomingGame method returns the expected value
     */
    it('#getUpcomingGames() should return value from observable', (done: DoneFn) => {

        const expectedGameSlug:string = 'upcoming game name';
    
            httpClientSpy.get.and.returnValue(of([{
                slug: expectedGameSlug
            }]));
    
            service.getUpcomingGames().subscribe(value => {
                expect(value[0].slug)
                    .withContext('expected game slug')
                    .toEqual(expectedGameSlug);
                done();
            });
            
            expect(httpClientSpy.get.calls.count())
                .withContext('one call')
                .toBe(1);
    })

    /**
     * Test that the getTopRatedGames method returns the expected value
     */
    it('#getTopRatedGames() should return value from observable', (done: DoneFn) => {

        const expectedGameSlug:string = 'most top-rated game name';
    
            httpClientSpy.get.and.returnValue(of([{
                slug: expectedGameSlug
            }]));
    
            service.getTopRatedGames().subscribe(value => {
                expect(value[0].slug)
                    .withContext('expected game slug')
                    .toEqual(expectedGameSlug);
                done();
            });
            
            expect(httpClientSpy.get.calls.count())
                .withContext('one call')
                .toBe(1);
    })

    /**
     * Test that the getFeaturedGame method returns the expected value
     */
    it('#getFeaturedGame() should return value from observable', (done: DoneFn) => {

        const expectedGameSlug:string = 'featured game';
    
            httpClientSpy.get.and.returnValue(of([{
                slug: expectedGameSlug
            }]));
    
            service.getFeaturedGame().subscribe(value => {
                expect(value[0].slug)
                    .withContext('expected game slug')
                    .toEqual(expectedGameSlug);
                done();
            });
            
            expect(httpClientSpy.get.calls.count())
                .withContext('one call')
                .toBe(1);
    })
});