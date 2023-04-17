import { ReviewService } from './review.service';
import { HttpClient } from '@angular/common/http';
import { of } from 'rxjs';

describe('ReviewService', () => {
  let service: ReviewService;
  let httpClientSpy: jasmine.SpyObj<HttpClient>;

  beforeEach(() => {
    httpClientSpy = jasmine.createSpyObj('HttpClient', ['get']);
    service = new ReviewService(httpClientSpy);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  /**
     * Test that the getRecentReviews method returns the expected value
     */
  it('#getRecentReviews() should return value from observable', (done: DoneFn) => {
            
    const expectedReviewDescription:string = 'This game sucks!';

    httpClientSpy.get.and.returnValue(of([{
        Description: expectedReviewDescription
    }]));

    service.getRecentReviews().subscribe(value => {
        expect(value[0].Description)
            .withContext('expected review description')
            .toEqual(expectedReviewDescription);
        done();
    });
    
    expect(httpClientSpy.get.calls.count())
        .withContext('one call')
        .toBe(1);
  });
});
